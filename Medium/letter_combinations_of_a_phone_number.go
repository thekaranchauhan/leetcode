package main

// Define the mapping of digits to letters as a global variable
var letterMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// Function to generate all possible letter combinations for a given string of digits
func letterCombinations(digits string) []string {
	// Initialize an empty result slice
	result := []string{}

	// Edge case: if the input string is empty, return an empty result
	if digits == "" {
		return result
	}

	// Initialize an empty slice to build each combination iteratively
	current := []byte{}

	// Helper function using backtracking to generate combinations
	var backtrack func(index int)
	backtrack = func(index int) {
		// Base case: if the current combination is complete (reached end of digits)
		if index == len(digits) {
			// Convert current byte slice to string and add to result
			result = append(result, string(current))
			return
		}

		// Get the current digit
		digit := digits[index]

		// Get the letters corresponding to the current digit from letterMap
		letters := letterMap[digit]

		// Iterate over each letter and explore the next digits recursively
		for _, letter := range letters {
			// Append the current letter to the current combination
			current = append(current, letter[0])
			// Recursively call backtrack for the next digit
			backtrack(index + 1)
			// Backtrack: remove the last added letter to try the next one
			current = current[:len(current)-1]
		}
	}

	// Start backtracking from the first digit
	backtrack(0)

	// Return the resulting combinations
	return result
}
