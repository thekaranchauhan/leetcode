package main

// strStr finds the first occurrence of needle in haystack using a simple sliding window approach.
func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	// Edge case: if needle is empty, return 0
	if nLen == 0 {
		return 0
	}

	// Loop through haystack to find the needle
	for i := 0; i <= hLen-nLen; i++ {
		// Check if substring of haystack matches the needle
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	// If needle is not found, return -1
	return -1
}
