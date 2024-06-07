package main

import (
	"strings"
)

// ReplaceWords replaces words in the sentence using the dictionary
func ReplaceWords(dictionary []string, sentence string) string {
	// Convert the dictionary into a set for quick lookup
	rootSet := make(map[string]struct{})
	for _, root := range dictionary {
		rootSet[root] = struct{}{}
	}

	// Split the sentence into words
	words := strings.Split(sentence, " ")

	// Function to find the shortest root for a given word
	findRoot := func(word string) string {
		for i := 1; i <= len(word); i++ {
			if _, exists := rootSet[word[:i]]; exists {
				return word[:i]
			}
		}
		return word
	}

	// Replace each word with the shortest matching root
	for i, word := range words {
		words[i] = findRoot(word)
	}

	// Join the words back into a sentence
	return strings.Join(words, " ")
}

// ! Time and Space Complexities are O(n*m) where n is the number of words in the sentence and m is the length of the longest word in the dictionary.
