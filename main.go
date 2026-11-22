package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)

		buf := make([]byte, 8)
		line := ""

		for {
			n, err := f.Read(buf)
			if err != nil {
				if len(line) != 0 {
					ch <- fmt.Sprintf("%s", line)
				}
				if errors.Is(err, io.EOF) {
					break
				}

				fmt.Printf("error: %s\n", err.Error())
				return
			}

			data := buf[:n]

			parts := strings.Split(string(data), "\n")
			parts_len := len(parts)

			for i, p := range parts {
				if i == (parts_len - 1) {
					line += p
					continue
				}
				ch <- fmt.Sprintf("%s%s", line, p)
				line = ""
			}

		}
	}()

	return ch
}

const inputFilePath = "./messages.txt"

func main() {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", inputFilePath, err)
	}
	defer file.Close()

	lines_ch := getLinesChannel(file)

	for line := range lines_ch {
		fmt.Printf("read: %s\n", line)
	}

}
