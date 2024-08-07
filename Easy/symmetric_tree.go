package main

// TreeNode represents a node in a binary tree.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// isSymmetric checks if the binary tree rooted at 'root' is symmetric.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true // An empty tree is symmetric.
	}
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two trees are mirror images of each other.
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true // Both trees are empty, so they are mirrors.
	}
	if t1 == nil || t2 == nil {
		return false // One tree is empty and the other is not, so they are not mirrors.
	}
	// Check if current nodes' values are equal and their subtrees are mirrors.
	return (t1.Val == t2.Val) &&
		isMirror(t1.Right, t2.Left) && // Right subtree of t1 is a mirror of the left subtree of t2.
		isMirror(t1.Left, t2.Right) // Left subtree of t1 is a mirror of the right subtree of t2.
}
