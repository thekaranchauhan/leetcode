package main

// specialArray finds the minimum value x such that there are exactly x numbers in nums that are greater than or equal to x.
func specialArray(nums []int) int {
	max := -1

	// Find the maximum value in nums
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	// Iterate over possible values of x
	for x := 1; x <= max; x++ {
		count := 0

		// Count numbers greater than or equal to x
		for _, num := range nums {
			if num >= x {
				count++
			}
		}

		// If count equals x, return x
		if count == x {
			return x
		}
	}

	return -1 // If no such x is found
}
