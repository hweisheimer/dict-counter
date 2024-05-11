package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func main() {
	instat, _ := os.Stdin.Stat()
	var wordBlob string

	if len(os.Args) > 1 {
		wordListBytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Could not read file", os.Args[1], err)
			os.Exit(1)
		}
		wordBlob = string(wordListBytes)
	} else if instat.Size() > 0 {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Could not read STDIN", err)
			os.Exit(1)
		}
		wordBlob = string(in)
	} else {
		fmt.Println("Expected filename or STDIN redirection (y pipes no work in golang?)")
		os.Exit(1)
	}

	wordBlob = normalize(wordBlob)
	charCounts := countCharacters(wordBlob)

	histogram := buildHistogram(charCounts, 50)
	if os.Getenv("lulz") != "" {
		histogram = strings.ReplaceAll(histogram, ">", "\U0001f95a")
	}
	fmt.Println(histogram)
}

func normalize(words string) string {
	words = strings.ToLower(words)

	// multi-line mode with ^$ mysteriously doesn't work well (only matches first line)
	// whitespace (\s) won't match file boundaries
	matcher := regexp.MustCompile(`\b.+'s\b`)

	// drops all possessive forms, because they are being used as separators
	// note that line endings and such are retained for the moment
	wordsSlice := matcher.Split(words, -1)
	words = strings.Join(wordsSlice, "")

	// could probably do this with utf8string package as well
	words = norm.NFD.String(words) // decompose unicode characters into base + combining marks
	wordRunes := []rune(words)
	wordRunes = slices.DeleteFunc(wordRunes, func(r rune) bool {
		// handily discards punctuation, diacritics, and misc junk that may be in the file
		return !unicode.In(r, unicode.Letter)
	})
	return string(wordRunes)
}

func countCharacters(words string) map[rune]uint32 {
	countMap := make(map[rune]uint32)
	for _, val := range words {
		countMap[val]++
	}
	return countMap
}

func buildHistogram(charCounts map[rune]uint32, chartWidth int) string {
	keys := []rune{}
	maxCount := uint32(0)
	for key, value := range charCounts {
		keys = append(keys, key)
		if value > maxCount {
			maxCount = value
		}
	}
	slices.Sort(keys)

	chart := ""
	for index, character := range keys {
		count := charCounts[character]
		renderWidth := int(float32(count) / float32(maxCount) * float32(chartWidth))
		chart += fmt.Sprintf("%s |%s (%d)", string(character), strings.Repeat(">", int(renderWidth)), count)
		if index+1 < len(keys) {
			chart += "\n"
		}
	}
	return chart
}
