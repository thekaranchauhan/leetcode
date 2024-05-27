package main

// longestCommonPrefix finds the longest common prefix string among an array of strings using binary search.
func longestCommonPrefix(strs []string) string {
	// If the input array is empty, return an empty string
	if len(strs) == 0 {
		return ""
	}

	// Find the minimum length of the strings in the array.
	// This helps to limit the search space for the longest common prefix.
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	// Initialize binary search boundaries
	low, high := 0, minLen

	// Perform binary search to find the longest common prefix length
	for low < high {
		// Calculate the middle point, favoring the upper half in case of tie
		mid := (low + high + 1) / 2
		// Check if the first `mid` characters are a common prefix
		if isCommonPrefix(strs, mid) {
			low = mid // Move the lower boundary up
		} else {
			high = mid - 1 // Move the upper boundary down
		}
	}

	// The longest common prefix is the substring of the first string up to `low` length
	return strs[0][:low]
}

// isCommonPrefix checks if the first `length` characters are the common prefix of all strings in the array.
func isCommonPrefix(strs []string, length int) bool {
	// Extract the prefix of the specified length from the first string
	prefix := strs[0][:length]
	// Compare the extracted prefix with the corresponding prefix of each string in the array
	for i := 1; i < len(strs); i++ {
		// If any string is shorter than `length` or does not match the prefix, return false
		if len(strs[i]) < length || strs[i][:length] != prefix {
			return false
		}
	}
	// If all strings match the prefix, return true
	return true
}
