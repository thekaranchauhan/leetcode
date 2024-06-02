package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Calculate the total length of the combined arrays
	totalLen := len(nums1) + len(nums2)
	// Calculate the index of the first median element
	firstMedianIdx := (totalLen - 1) / 2
	// Calculate the index of the second median element if the total length is even
	lastMedianIdx := firstMedianIdx + (1^totalLen)&1

	// Initialize a slice to store the merged elements up to the necessary indices
	merged := make([]int, 0, lastMedianIdx+1)

	// Pointers for iterating through nums1 and nums2
	i, j := 0, 0
	// Merge elements from nums1 and nums2 until the merged slice contains the necessary indices
	for i+j <= lastMedianIdx {
		if i < len(nums1) && (j >= len(nums2) || nums1[i] <= nums2[j]) {
			// If nums1[i] is less than or equal to nums2[j], append nums1[i] to the merged slice
			// or if nums2 is exhausted
			merged = append(merged, nums1[i])
			i++
		} else {
			// Otherwise, append nums2[j] to the merged slice
			merged = append(merged, nums2[j])
			j++
		}
	}

	// Calculate the median from the elements at firstMedianIdx and lastMedianIdx
	return float64(merged[firstMedianIdx]+merged[lastMedianIdx]) / 2.0
}

// ! Finding median of two sorted arrays in O(log(min(m, n))) time complexity.
