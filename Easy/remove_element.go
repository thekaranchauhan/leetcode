package main

// removeElement removes all occurrences of val in nums in-place and returns the number of elements not equal to val.
func removeElement(nums []int, val int) int {
	// Initialize a pointer `k` to track the position of the next non-val element.
	k := 0

	// Iterate through the array with the fast pointer `i`.
	for i := 0; i < len(nums); i++ {
		// If the current element is not equal to val, it should be kept.
		if nums[i] != val {
			// Move the current element to the position indicated by `k`.
			nums[k] = nums[i]
			k++
		}
	}

	// `k` represents the number of elements not equal to val.
	return k
}
