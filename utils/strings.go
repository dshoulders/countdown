package utils

import "strings"

func RemoveChar(str string, character rune) (string, bool) {
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

func GetFirstRune(str string) rune {
	var first rune
	for _, char := range str {
		first = char
		break
	}
	return first
}
