package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
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

		fmt.Print("#", line)
	}
}
