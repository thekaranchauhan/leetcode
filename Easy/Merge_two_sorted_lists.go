package main

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists and returns the head of the merged list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to serve as the starting point
	dummy := &ListNode{}
	current := dummy

	// Iterate through both lists until one of them is exhausted
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach the remaining part of the list that is not yet exhausted
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the merged list starting from the next node of the dummy
	return dummy.Next
}

// ! This solution efficiently merges two sorted linked lists with a time complexity of O(n + m), where n and m are the lengths of list1 and list2, respectively. The space complexity is O(1) as the merging is done in place.
