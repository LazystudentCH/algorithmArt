package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else {
		if root.Left == nil {
			r := root.Right
			root.Right = nil
			return r
		}

		if root.Right == nil {
			l := root.Left
			root.Left = nil
			return l
		}
		min := findmin(root.Right)
		min.Right = removemin(root.Right)
		min.Left = root.Left
		root.Left = nil
		root.Right = nil

		return min
	}

}

func findmin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil {
		return root
	}

	return findmin(root.Left)
}

func removemin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil {
		r := root.Right
		return r
	}

	root.Left = removemin(root.Left)
	return root

}
