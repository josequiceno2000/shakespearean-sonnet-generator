package main

import (
	"unicode"
)

func countSyllables(pronunciation string) int {
	count := 0
	for _, char := range pronunciation {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}