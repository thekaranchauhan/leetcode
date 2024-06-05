package main

import (
	"math"
)

func commonChars(words []string) []string {
	// Initialize an array to keep the minimum frequency of each character
	// Set initial frequency to a large number to find minimum across all words
	minFreq := make([]int, 26)
	for i := range minFreq {
		minFreq[i] = math.MaxInt32
	}

	// Iterate over each word in the input array
	for _, word := range words {
		// Create a frequency count for the current word
		freq := make([]int, 26)
		for _, char := range word {
			freq[char-'a']++
		}

		// Update the global minimum frequency for each character
		for i := 0; i < 26; i++ {
			if freq[i] < minFreq[i] {
				minFreq[i] = freq[i]
			}
		}
	}

	// Prepare the result array with common characters based on minimum frequency
	result := []string{}
	for i := 0; i < 26; i++ {
		for minFreq[i] > 0 {
			result = append(result, string('a'+i))
			minFreq[i]--
		}
	}

	return result
}

// ! This solution has a time complexity of O(n * m) where n is the number of words in the input array and m is the average length of each word.
