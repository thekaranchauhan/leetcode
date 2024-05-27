package main

// isValid checks if the input string s is a valid sequence of parentheses.
func isValid(s string) bool {
	// Map to hold matching pairs of parentheses
	pairMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Stack to keep track of opening brackets
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		// If the character is a closing bracket
		if match, ok := pairMap[char]; ok {
			// Check if the stack is empty or the top of the stack does not match
			if len(stack) == 0 || stack[len(stack)-1] != match {
				return false
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		} else {
			// Push the opening bracket onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets are matched
	return len(stack) == 0
}

// ! This solution is optimal with a time complexity of O(n), where n is the length of the string, because each character is processed once. The space complexity is O(n) in the worst case, where all characters are opening brackets.
