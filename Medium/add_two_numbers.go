package main

// ListNode represents a node in the linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists and returns the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize a dummy head node for the result linked list
	dummy := &ListNode{}
	curr := dummy
	carry := 0 // Variable to store carry-over

	// Iterate through both linked lists simultaneously
	for l1 != nil || l2 != nil || carry > 0 {
		// Add values of current nodes and carry-over
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		// Create a new node with the sum of digits
		curr.Next = &ListNode{Val: carry % 10}
		carry /= 10 // Update carry-over
		curr = curr.Next
	}

	return dummy.Next // Return the result linked list (excluding the dummy head)
}
