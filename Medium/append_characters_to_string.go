package main

// Function to calculate the minimum number of characters to append
// to s so that t becomes a subsequence of s
func minCharactersToAdd(s string, t string) int {
	n, m := len(s), len(t) // Get the lengths of both strings
	j := 0                 // Pointer for iterating over string t

	// Traverse string s from left to right
	for i := 0; i < n; i++ {
		// If the current character of s matches the current character of t
		if j < m && s[i] == t[j] {
			j++ // Move to the next character in t
		}
	}

	// The number of characters to append is the remaining characters in t
	return m - j
}

// ! Time complexity: O(n+m), where n and m are the lengths of strings s and t, respectively.
