package main

import (
	"math"
	"unicode"
)

// myAtoi converts a string to a 32-bit signed integer.
func myAtoi(s string) int {
	n := len(s) // Get the length of the string
	i := 0      // Index to traverse the string

	// Step 1: Ignore leading whitespaces
	for i < n && unicode.IsSpace(rune(s[i])) {
		i++
	}

	// Step 2: Check for sign
	sign := 1 // Default sign is positive
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1 // Change sign to negative if '-' is found
		}
		i++ // Move to the next character after the sign
	}

	// Step 3: Convert digits to integer
	num := 0 // Initialize the result
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0') // Convert the character to its integer value

		// Step 4: Check for overflow and clamp to 32-bit signed integer range
		if num > (math.MaxInt32-digit)/10 {
			// If adding the new digit causes overflow, return the appropriate boundary
			if sign == 1 {
				return math.MaxInt32 // Positive overflow
			} else {
				return math.MinInt32 // Negative overflow
			}
		}

		num = num*10 + digit // Add the digit to the result
		i++                  // Move to the next character
	}

	return num * sign // Apply the sign to the result
}

// ! Time complexity: O(n), where n is the length of the input string s.
