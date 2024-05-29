package main

// longestPalindrome finds the longest palindromic substring in the given string.
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0 // Initialize the start and end indices of the longest palindrome

	// Helper function to expand around the center
	expandAroundCenter := func(left, right int) {
		// Expand while the characters are equal and indices are within bounds
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		left++  // Move back to the last valid position
		right-- // Move back to the last valid position
		// Update start and end if a longer palindrome is found
		if right-left > end-start {
			start, end = left, right
		}
	}

	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)   // Check for odd-length palindromes
		expandAroundCenter(i, i+1) // Check for even-length palindromes
	}

	// Return the longest palindromic substring
	return s[start : end+1]
}

// ! This solution has a time complexity of O(n^2) and a space complexity of O(1).
