package main

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLength := 0                     // To keep track of the maximum length of substring found
	start := 0                         // The starting index of the current substring window
	charIndexMap := make(map[byte]int) // A map to store the last seen index of each character

	for end := 0; end < n; end++ {
		if index, found := charIndexMap[s[end]]; found && index >= start {
			// If the character is found in the map and its last seen index is within the current window
			start = index + 1 // Move the start to the right of the last seen index
		}
		charIndexMap[s[end]] = end // Update the last seen index of the current character

		// Calculate the current length of the substring and update maxLength if necessary
		if currentLength := end - start + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength // Return the maximum length of the substring found
}

// ! Solution with with a time complexity of O(n) and space complexity of O(min(n, m)), where m is the size of the character set (which can be considered constant).
