package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("some_file.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 16)

	for {
		n, err := reader.Read(buf)

		if err != nil {

			if err != io.EOF {

				log.Fatal(err)
			}

			break
		}

		fmt.Print(string(buf[0:n]))
	}

	fmt.Println()
}
