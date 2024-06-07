package main

func plusOne(digits []int) []int {
	n := len(digits)

	// Iterate from the end of the digits array
	for i := n - 1; i >= 0; i-- {
		// If the digit is less than 9, just increment and return
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// Set the current digit to 0 if it was 9 and continue the loop
		digits[i] = 0
	}

	// If we are here it means all the digits were 9 and we need an extra digit at the beginning
	return append([]int{1}, digits...)
}

// ! Time and Space Complexities are O(n) where n is the number of digits in the input array.
