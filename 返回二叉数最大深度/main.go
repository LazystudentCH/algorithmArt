package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	result := help(root, 0)
	return result
}

func help(root *TreeNode, deepth int) int {
	if root == nil {
		return deepth
	}

	deepth++

	left_depth := help(root.Left, deepth)
	right_depth := help(root.Right, deepth)

	return max(left_depth, right_depth)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
