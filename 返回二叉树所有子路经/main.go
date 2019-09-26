package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		return append(res, strconv.Itoa(root.Val))
	}

	rl := binaryTreePaths(root.Left)
	for i := 0; i < len(rl); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+rl[i])
	}

	rr := binaryTreePaths(root.Right)
	for i := 0; i < len(rr); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+rr[i])
	}

	return res
}
