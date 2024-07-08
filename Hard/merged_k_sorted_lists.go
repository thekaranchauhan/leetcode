package main

import (
	"container/heap"
)

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeHeap is a min-heap for ListNode pointers.
type ListNodeHeap []*ListNode

// Implementing heap.Interface methods for ListNodeHeap.
func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an item (ListNode) to the heap.
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Pop removes and returns the smallest item (ListNode) from the heap.
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// mergeKLists merges k sorted linked lists into one sorted linked list.
func mergeKLists(lists []*ListNode) *ListNode {
	// Initialize the min-heap
	h := &ListNodeHeap{}
	heap.Init(h)

	// Push the head of each list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	// Dummy node to simplify the result list handling
	dummy := &ListNode{}
	tail := dummy

	// Merge process
	for h.Len() > 0 {
		// Pop the smallest element from the heap
		minNode := heap.Pop(h).(*ListNode)
		tail.Next = minNode
		tail = tail.Next

		// Push the next element of the popped list back into the heap if exists
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	// Return the merged sorted list
	return dummy.Next
}
