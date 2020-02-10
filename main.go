package main

import (
	"countdown/dictionary"
	"countdown/utils"
	"fmt"
	"os"
	"runtime"
	"time"

	flag "github.com/ogier/pflag"
)

var (
	letters string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	chunkCount := runtime.NumCPU()

	words, err := dictionary.ReadLines("./dictionary/words.txt")
	if err != nil {
		fmt.Println("readlines: %s", err)
	}

	c := make(chan []string, chunkCount)

	start := time.Now()

	wordChunks := utils.ChunkArray(chunkCount, words)
	var winners []string

	for _, chunk := range wordChunks {
		go func(chunk []string) {
			c <- getMatches(chunk, letters)
		}(chunk)
	}

	for i := 0; i < chunkCount; i++ {
		winners = append(winners, <-c...)
	}

	longest := utils.GetLongest(winners)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(len(longest), winners, longest)
	fmt.Println(elapsed)
}

func getMatches(words []string, letters string) []string {
	var matches []string

	for _, word := range words {

		if len(word) > len(letters) {
			continue
		}

		if checkForMatch(word, letters) {
			matches = append(matches, word)
		}
	}

	return matches
}

func checkForMatch(word string, letters string) bool {

	firstLetter := utils.GetFirstRune(word)
	letters, found := utils.RemoveChar(letters, firstLetter)

	if found {
		word, _ = utils.RemoveChar(word, firstLetter)
	}

	if len(word) == 0 {
		return true
	}

	if found {
		return checkForMatch(word, letters)
	}

	return false
}

func init() {
	flag.StringVarP(&letters, "letters", "l", "", "Letter Palette")
}
