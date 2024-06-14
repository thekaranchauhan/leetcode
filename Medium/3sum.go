package main

import (
	"sort"
)

// threeSum finds all unique triplets in the array that sum up to zero.
func threeSum(nums []int) [][]int {
	// Sort the array to facilitate the two-pointer approach.
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	// Iterate through the array, considering each element as a potential first element of a triplet.
	for i := 0; i < n-2; i++ {
		// Skip duplicate elements to avoid duplicate triplets in the result.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Initialize two pointers: one just after the current element, and the other at the end of the array.
		left, right := i+1, n-1

		// Use the two-pointer technique to find pairs that, with nums[i], sum to zero.
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// If the sum is zero, we found a valid triplet.
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Move both pointers inward and skip duplicates.
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				// If the sum is less than zero, move the left pointer to the right to increase the sum.
				left++
			} else {
				// If the sum is greater than zero, move the right pointer to the left to decrease the sum.
				right--
			}
		}
	}

	return result
}
