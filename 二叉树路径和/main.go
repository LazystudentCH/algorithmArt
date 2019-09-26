package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := helper(root, sum)
	res += pathSum(root.Left, sum)
	res += pathSum(root.Right, sum)

	return res

}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res = res + 1
	}

	res += helper(root.Left, sum-root.Val)
	res += helper(root.Right, sum-root.Val)

	return res
}
