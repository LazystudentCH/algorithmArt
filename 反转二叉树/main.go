package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	t := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(t)

	return root
}
