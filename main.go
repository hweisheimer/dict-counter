package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	instat, _ := os.Stdin.Stat()
	var wordList string

	if len(os.Args) > 1 {
		wordListBytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Could not read file", os.Args[1], err)
			os.Exit(1)
		}
		wordList = string(wordListBytes)
	} else if instat.Size() > 0 {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Could not read STDIN", err)
			os.Exit(1)
		}
		wordList = string(in)
	} else {
		fmt.Println("Expected filename or STDIN redirection (y pipes no work in golang?)")
		os.Exit(1)
	}

	fmt.Println("Received word list contents:\n", wordList)
}

func Normalize(words string) string {
	return "lol"
}
