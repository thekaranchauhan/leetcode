package main

func longestPalindrome(s string) int {
	// Create a map to count the occurrences of each character
	charCount := make(map[rune]int)

	// Iterate over each character in the string and count occurrences
	for _, char := range s {
		charCount[char]++
	}

	// Initialize length of the longest palindrome
	length := 0
	// Flag to check if there's any character with an odd count
	oddFound := false

	// Iterate over character counts
	for _, count := range charCount {
		if count%2 == 0 {
			// If the count is even, it can fully contribute to the palindrome
			length += count
		} else {
			// If the count is odd, all but one can contribute to the palindrome
			length += count - 1
			oddFound = true
		}
	}

	// If there was any character with an odd count, add one to the length
	// to place exactly one such character in the middle of the palindrome
	if oddFound {
		length++
	}

	return length
}

// ! This solution has a time complexity of O(n) where n is the length of the input string s.
