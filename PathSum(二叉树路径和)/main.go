package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Right == nil && root.Left == nil {
		return true
	}

	if root.Right == nil && root.Left == nil && root.Val != sum {
		return false
	}

	sum = sum - root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
