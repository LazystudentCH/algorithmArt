package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return helper(root, root.Val)
}

func helper(root *TreeNode, Val int) bool {
	if root == nil {
		return true
	}

	if root.Val != Val {
		return false
	}

	return helper(root.Left, Val) && helper(root.Right, Val)
}
