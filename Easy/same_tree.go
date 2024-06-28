package main

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree checks if two binary trees are the same
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base cases
	if p == nil && q == nil {
		return true // Both nodes are nil
	}
	if p == nil || q == nil {
		return false // One node is nil, the other is not
	}
	if p.Val != q.Val {
		return false // Nodes have different values
	}

	// Recursive cases: check left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
