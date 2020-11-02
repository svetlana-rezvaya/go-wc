package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func countWords(text string) int {
	spaceCount := 0
	for _, character := range text {
		if character == ' ' {
			spaceCount = spaceCount + 1
		}
	}

	wordCount := spaceCount + 1
	return wordCount
}

func main() {
	lineCount := 0
	wordCount := 0
	byteCount := 0
	characterCount := 0
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
		wordCount = wordCount + countWords(line)
		byteCount = byteCount + len(line)
		characterCount = characterCount + utf8.RuneCountInString(line)

		fmt.Print("#", line)
	}

	fmt.Println("amount of lines =", lineCount)
	fmt.Println("amount of words =", wordCount)
	fmt.Println("amount of bytes =", byteCount)
	fmt.Println("amount of characters =", characterCount)
}
