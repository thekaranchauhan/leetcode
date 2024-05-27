package main

// twoSum finds two numbers in nums that add up to target and returns their indices.
// It assumes that there is exactly one solution and does not use the same element twice.
func twoSum(nums []int, target int) []int {
	// Initialize a map to store numbers and their indices.
	// Preallocating the map with the length of nums to avoid resizing.
	numToIndex := make(map[int]int, len(nums))

	// Iterate through the list of numbers with their indices.
	for i, num := range nums {
		// Calculate the complement of the current number.
		// The complement is the number that we need to find in the map.
		complement := target - num

		// Check if the complement exists in the map.
		// If found, return the indices of the complement and the current number.
		if index, found := numToIndex[complement]; found {
			return []int{index, i}
		}

		// Add the current number and its index to the map.
		// This step is done after the complement check to avoid using the same element twice.
		numToIndex[num] = i
	}

	// Return nil if no solution is found.
	// This line should never be reached because the problem guarantees exactly one solution.
	return nil
}
