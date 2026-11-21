package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 8)
	line := ""

	for {
		n, err := file.Read(buf)

		data := buf[:n]

		parts := strings.Split(string(data), "\n")
		parts_len := len(parts)

		for i, p := range parts {
			if i == (parts_len - 1) {
				line += p
				continue
			}
			fmt.Printf("read: %s%s\n", line, p)
			line = ""
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(line) != 0 {
					fmt.Printf("read: %s\n", line)
				}
				break
			}
			log.Fatal(err)
		}
	}
}
