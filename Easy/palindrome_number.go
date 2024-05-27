package main

func isPalindrome(x int) bool {
	// Handle negative numbers and numbers ending with zero
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// Reverse half of the number and compare with the other half
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// If the length of the number is odd, we can ignore the middle digit
	// For example, if the input is 12321, after the loop, reversed will be 123,
	// and x will be 12. So we need to divide reversed by 10 to remove the middle digit.
	return x == reversed || x == reversed/10
}
