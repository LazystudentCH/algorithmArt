package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	l := rangeSumBST(root.Left, L, R)
	r := rangeSumBST(root.Right, L, R)
	if root.Val >= L && root.Val <= R {
		return root.Val + l + r
	}

	return l + r
}
