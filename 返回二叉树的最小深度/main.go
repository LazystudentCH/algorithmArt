package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root != nil {
		if root.Left != nil && root.Right != nil {
			return 1 + min(minDepth(root.Left), minDepth(root.Right))
		} else if root.Left != nil {
			return 1 + minDepth(root.Left)
		} else if root.Right != nil {
			return 1 + minDepth(root.Right)
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	//[3,9,20,null,null,15,7]
	r := new(TreeNode)
	r.Val = 3
	r.Left = &TreeNode{
		9,
		&TreeNode{
			0,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	fmt.Println(minDepth(r))
}
