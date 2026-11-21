package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 8)

	for {
		len, err := file.Read(buf)

		data := buf[:len]
		fmt.Printf("read: %s\n", data)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
	}
}
