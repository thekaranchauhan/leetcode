package main

// searchInsert finds the index of the target in the sorted array or the index where it should be inserted.
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1 // Initialize left and right pointers

	for left <= right {
		// Calculate the middle index
		mid := left + (right-left)/2

		if nums[mid] == target {
			// If target is found at mid, return mid
			return mid
		} else if nums[mid] < target {
			// If target is greater than the mid element, search in the right half
			left = mid + 1
		} else {
			// If target is less than the mid element, search in the left half
			right = mid - 1
		}
	}

	// If target is not found, left will be the index where it should be inserted
	return left
}
