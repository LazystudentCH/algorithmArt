package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	num1 := nums[:mid]
	num2 := nums[mid+1:]
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBST(num1)
	root.Right = sortedArrayToBST(num2)

	return root
}
