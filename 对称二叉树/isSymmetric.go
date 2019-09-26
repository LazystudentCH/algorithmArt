package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if !(root.Left.Val == root.Right.Val) {
		return false
	}

	return isSymmetric(root.Left) && isSymmetric(root.Right)
}
