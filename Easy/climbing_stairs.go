package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// Initialize the base cases
	first, second := 1, 2

	// Iterate from step 3 to n
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}
