package main

// removeDuplicates removes duplicates in-place from a sorted array and returns the number of unique elements.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// `k` is the slow pointer that marks the position of the last unique element.
	k := 0

	// Iterate through the array starting from the second element.
	for i := 1; i < len(nums); i++ {
		// If the current element is different from the last unique element, it's a new unique element.
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	// The number of unique elements is `k + 1`.
	return k + 1
}

// ! This solution has a time complexity of O(n), where n is the length of the array, as each element is processed once. The space complexity is O(1) because the algorithm uses a constant amount of extra space.
