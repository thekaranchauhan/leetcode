package main

import "strings"

// convert takes a string s and converts it into a zigzag pattern with the given number of rows.
func convert(s string, numRows int) string {
	// Edge case: if numRows is 1 or greater than or equal to the length of s,
	// the zigzag pattern is just the string itself.
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Create an array to hold strings for each row
	rows := make([]string, numRows)
	currRow := 0       // The current row index we are placing characters into
	goingDown := false // A flag to indicate the direction of traversal through the rows

	// Distribute characters to the rows in a zigzag pattern
	for _, char := range s {
		// Append the current character to the current row
		rows[currRow] += string(char)
		// Change direction when reaching the first or last row
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}
		// Move to the next row based on the direction
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	// Concatenate all rows to get the final string
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row)
	}

	return result.String() // Return the final zigzag converted string
}

// ! The time complexity of this solution is O(n), where n is the length of the input string s and the space complexity is O(n).
