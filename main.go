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
	// Parse command line params
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	chunkCount := runtime.NumCPU()

	// Read all words in the dictionary file
	words, err := dictionary.ReadLines("./dictionary/words.txt")
	if err != nil {
		fmt.Println("readlines: %s", err)
	}

	c := make(chan []string, chunkCount)

	// Begin performance test
	start := time.Now()

	// Split the words into chunks to send to goroutines
	wordChunks := utils.ChunkSlice(chunkCount, words)
	var winners []string

	// Find all matches in each chunk
	for _, chunk := range wordChunks {
		go func(chunk []string) {
			c <- getMatches(chunk, letters)
		}(chunk)
	}

	// Collect and concat all matches into single slice
	for i := 0; i < chunkCount; i++ {
		winners = append(winners, <-c...)
	}

	// Find the longest matched word
	longest := utils.GetLongest(winners)

	// End performance test
	t := time.Now()
	elapsed := t.Sub(start)

	// Print results
	fmt.Println(len(longest), winners, longest)
	fmt.Println(elapsed)
}

func getMatches(words []string, letters string) []string {
	var matches []string

	// Check each word in words to see if it can be matched
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

	// Check first letter of word exists in the letters palette and remove it to prevent it being matched a again
	firstLetter := utils.GetFirstRune(word)
	letters, found := utils.RemoveRune(letters, firstLetter)

	if found {
		// Remove it from the word so we know what characters we still need to check
		word, _ = utils.RemoveRune(word, firstLetter)
	}

	if len(word) == 0 {
		// All letters have been matched in the word
		return true
	}

	if found {
		// Recursively continue checking the rest of the word
		return checkForMatch(word, letters)
	}

	// A letter in the word has not been found - no match
	return false
}

func init() {
	flag.StringVarP(&letters, "letters", "l", "", "Letter Palette")
}
