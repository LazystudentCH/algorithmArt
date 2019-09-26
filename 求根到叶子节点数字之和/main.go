package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	result := 0
	res := helper(root)
	for i := 0; i < len(res); i++ {
		k, _ := strconv.Atoi(res[i])
		result += k
	}
	return result
}

func helper(root *TreeNode) []string {
	res := []string{}

	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		return append(res, strconv.Itoa(root.Val))
	}

	l := helper(root.Left)
	for i := 0; i < len(l); i++ {
		res = append(res, strconv.Itoa(root.Val)+l[i])
	}
	r := helper(root.Right)
	for i := 0; i < len(r); i++ {
		res = append(res, strconv.Itoa(root.Val)+r[i])
	}

	return res

}
