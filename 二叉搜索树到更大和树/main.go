package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	helper(root, &sum)
	return root
}

func helper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	helper(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	helper(root.Left, sum)
}
