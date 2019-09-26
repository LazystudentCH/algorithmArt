package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ka int
var res int

func kthSmallest(root *TreeNode, k int) int {
	ka = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	ka--
	if ka == 0 {
		res = root.Val
		return
	}
	dfs(root.Right)
}
