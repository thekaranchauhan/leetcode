package main

import (
	"math"
)

// reverse takes a signed 32-bit integer and returns its digits reversed.
// If the reversed integer overflows, it returns 0.
func reverse(x int) int {
	const intMax = math.MaxInt32 // Maximum 32-bit integer value (2147483647)
	const intMin = math.MinInt32 // Minimum 32-bit integer value (-2147483648)

	result := 0 // Initialize result to 0
	for x != 0 {
		pop := x % 10 // Get the last digit of x
		x /= 10       // Remove the last digit from x

		// Check for overflow before adding the digit to the result
		if result > (intMax-pop)/10 {
			return 0 // Return 0 if adding pop would cause overflow
		}
		// Check for underflow before adding the digit to the result
		if result < (intMin-pop)/10 {
			return 0 // Return 0 if adding pop would cause underflow
		}

		result = result*10 + pop // Append the digit to the result
	}

	return result // Return the reversed integer
}

// ! The time complexity of this solution is O(log(x)), where x is the input integer, and the space complexity is O(1).
