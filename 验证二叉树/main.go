package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := helper(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}

	return true

}

func helper(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	l := helper(root.Left)
	r := helper(root.Right)

	res = append(l, root.Val)
	res = append(res, r...)

	return res
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	head, tail := 0, len(nums)-1
	x, i := nums[0], 1

	for head < tail {
		if x < nums[i] {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}

	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}
