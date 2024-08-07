package main

// generateParenthesis generates all combinations of well-formed parentheses for a given number of pairs 'n'.
func generateParenthesis(n int) []string {
	var result []string

	// backtrack is a recursive helper function that builds the combinations.
	// 'current' is the current string being built,
	// 'open' is the count of open parentheses used,
	// 'close' is the count of close parentheses used.
	var backtrack func(current string, open, close int)

	backtrack = func(current string, open, close int) {
		// If the current string has reached the maximum length (2 * n), add it to the result list.
		if len(current) == n*2 {
			result = append(result, current)
			return
		}
		// If the number of open parentheses used is less than 'n', add an open parenthesis and recurse.
		if open < n {
			backtrack(current+"(", open+1, close)
		}
		// If the number of close parentheses used is less than the number of open parentheses, add a close parenthesis and recurse.
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	// Start the backtracking process with an empty string and zero open/close parentheses.
	backtrack("", 0, 0)
	return result
}
