package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
	"unicode"
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

	wordList = normalize(wordList)
	// TODO: the things
}

func normalize(words string) string {
	words = strings.ToLower(words)

	// multi-line mode with ^$ mysteriously doesn't work well (only matches first line)
	// whitespace (\s) won't match file boundaries
	matcher := regexp.MustCompile(`\b.+'s\b`)

	wordsSlice := matcher.Split(words, -1)
	words = strings.Join(wordsSlice, "")
	cleanWordRunes := slices.DeleteFunc([]rune(words), func(r rune) bool {
		// would be more comprehensive to remove anything that's not alpha, but then I'd need another test ;)
		return unicode.In(r, unicode.Punct, unicode.White_Space)
	})

	return string(cleanWordRunes)
}
