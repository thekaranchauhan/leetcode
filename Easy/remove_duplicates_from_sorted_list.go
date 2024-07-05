package main

// ListNode defines a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates removes duplicates from a sorted linked list.
func deleteDuplicates(head *ListNode) *ListNode {
	// If the list is empty, return nil
	if head == nil {
		return nil
	}

	// Initialize a pointer to traverse the list
	current := head

	// Traverse the list until the end
	for current != nil && current.Next != nil {
		// If the current node's value is the same as the next node's value
		if current.Val == current.Next.Val {
			// Skip the next node by linking to the node after next
			current.Next = current.Next.Next
		} else {
			// Move to the next node
			current = current.Next
		}
	}

	// Return the modified list with duplicates removed
	return head
}
