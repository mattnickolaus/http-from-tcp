package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(conn net.Conn) <-chan string {
	ch := make(chan string)

	go func() {
		defer conn.Close()
		defer close(ch)

		buf := make([]byte, 8)
		line := ""

		for {
			n, err := conn.Read(buf)
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

// const inputFilePath = "./messages.txt"
const port = ":42069"

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could listen on port %s: %s\n", port, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Could accept connection: %s\n", err)
		}
		fmt.Println("Connection Accepted")

		lines_ch := getLinesChannel(conn)
		for line := range lines_ch {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection has been closed")
	}

}
