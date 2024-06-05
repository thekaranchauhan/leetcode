package main

func isMatch(s string, p string) bool {
	// Get the lengths of the input string and the pattern
	m, n := len(s), len(p)

	// Create a 2D dp array with default false values
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string matches empty pattern
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* which can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// If characters match or pattern has '.', carry over the match status
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// If pattern has '*', check zero occurrences or more occurrences of the preceding character
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.'))
			}
		}
	}

	// The answer is in the bottom-right cell of the dp table
	return dp[m][n]
}

// ! Time and Space Complexities are O(m*n) where m is the length of the input string and n is the length of the pattern.
