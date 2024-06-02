package main

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		// Swap the characters at left and right indices
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// ! Reversing a string represented as an array of characters in-place using O(1) extra memory.
