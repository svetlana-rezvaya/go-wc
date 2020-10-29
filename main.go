package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lineCount := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			// err == io.EOF
			if line == "" {
				break
			}
		}

		lineCount = lineCount + 1

		fmt.Print("#", line)
	}

	fmt.Println("amount of lines =", lineCount)
}
