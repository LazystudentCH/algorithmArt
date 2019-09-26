package main

import (
	"fmt"
)

/*
输入一棵二叉树前序遍历和中序遍历的结果，请重建该二叉树。

注意:

二叉树中每个节点的值都互不相同；
输入的前序遍历和中序遍历一定合法；
样例
给定：
前序遍历是：[3, 9, 20, 15, 7]
中序遍历是：[9, 3, 15, 20, 7]

返回：[3, 9, 20, null, null, 15, 7, null, null, null, null]
返回的二叉树如下所示：
    3
   / \
  9  20
    /  \
   15   7
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func main() {

}
