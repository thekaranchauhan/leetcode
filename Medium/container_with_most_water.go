package main

// maxArea calculates the maximum area of water that can be contained
// between two lines on the x-axis represented by the heights in the input slice.
func maxArea(height []int) int {
	// Initialize two pointers, one starting at the beginning (left)
	// and one at the end (right) of the height array.
	left, right := 0, len(height)-1
	// Initialize a variable to keep track of the maximum area found.
	maxArea := 0

	// Loop until the two pointers meet.
	for left < right {
		// Calculate the width of the current container.
		width := right - left
		// Determine the height of the container, which is the minimum of
		// the heights at the left and right pointers.
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++ // Move the left pointer to the right.
		} else {
			h = height[right]
			right-- // Move the right pointer to the left.
		}
		// Calculate the area of the current container.
		area := width * h
		// Update maxArea if the current area is greater than the previous maxArea.
		if area > maxArea {
			maxArea = area
		}
	}
	// Return the maximum area found.
	return maxArea
}

// ! Finding maximum area in O(n) time complexity and O(1) space complexity.
