package utils

import "strings"

// RemoveRune - Removes a character from a string if found and returns the string and whether it was found
func RemoveRune(str string, character rune) (string, bool) {
	found := false

	filter := func(r rune) rune {

		// We only want to remove one character at a time
		if found {
			return r
		}

		if r == character {
			found = true
			return -1
		}
		return r
	}

	return strings.Map(filter, str), found
}

// GetFirstRune - Returns first rune of a srting
func GetFirstRune(str string) rune {
	var first rune
	for _, char := range str {
		first = char
		break
	}
	return first
}
