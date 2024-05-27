package main

import (
	"fmt"
)

func romanToInt(s string) int {
	// Initialize result to store the final integer value.
	result := 0

	// Initialize prevValue to keep track of the previous Roman numeral's integer value.
	prevValue := 0

	// Iterate over the characters of the Roman numeral string from right to left.
	for i := len(s) - 1; i >= 0; i-- {
		// Get the integer value of the current Roman numeral character.
		value := getValue(s[i])

		// Compare the current value with the previous value to determine whether to add or subtract.
		if value < prevValue {
			result -= value // Subtract if current value is less than previous value.
		} else {
			result += value // Add if current value is greater than or equal to previous value.
		}

		// Update prevValue for the next iteration.
		prevValue = value
	}

	// Return the computed integer value of the Roman numeral string.
	return result
}

// getValue returns the integer value of a single Roman numeral character.
func getValue(ch byte) int {
	// Switch statement to match the character with its corresponding integer value.
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func main() {
	// Test cases
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
}
