package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	res := 0
	d := -1
	helper(root, 0, &d, &res)
	return res

}

func helper(root *TreeNode, l int, d, res *int) {
	if root == nil {
		return
	}
	if l > *d {
		*d = l
		*res = root.Val
	}

	helper(root.Left, l+1, d, res)
	helper(root.Right, l+1, d, res)
}
