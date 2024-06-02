package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	// Trim trailing spaces
	s = strings.TrimSpace(s)
	// Split the string by spaces
	words := strings.Split(s, " ")
	// Get the last word
	lastWord := words[len(words)-1]
	// Return the length of the last word
	return len(lastWord)
}
