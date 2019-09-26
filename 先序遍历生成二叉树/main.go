package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	index := 0
	for i := 0; i < len(preorder); i++ {
		if preorder[i] > root.Val {
			index = i
			break
		}
	}
	if index != 0 {
		root.Left = bstFromPreorder(preorder[1:index])
		root.Right = bstFromPreorder(preorder[index:])
	} else {
		root.Left = bstFromPreorder(preorder[1:])
		root.Right = nil
	}

	return root
}
