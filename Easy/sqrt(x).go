package main

func mySqrt(x int) int {
	// Special case: if x is 0, the square root is 0
	if x == 0 {
		return 0
	}

	// Initialize the binary search boundaries
	left, right := 1, x
	result := 0

	// Perform binary search
	for left <= right {
		// Calculate the middle point to avoid overflow
		mid := left + (right-left)/2

		// If mid*mid <= x, mid is a candidate for the square root
		if mid <= x/mid {
			result = mid   // Update result
			left = mid + 1 // Move left boundary to search higher values
		} else {
			// If mid*mid > x, discard mid and search lower values
			right = mid - 1
		}
	}

	// When the loop ends, result holds the largest integer whose square is <= x
	return result
}
