package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := CountDepth(root.Left)
	r := CountDepth(root.Right)

	if max(l, r)-min(l, r) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func CountDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(CountDepth(root.Left), CountDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
