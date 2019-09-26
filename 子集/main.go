package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	helper(nums, []int{}, &res)
	return append(res, []int{})
}

func helper(nums, temp []int, res *[][]int) {
	if len(temp) != 0 {
		a := make([]int, len(temp))
		copy(a, temp)
		*res = append(*res, a)
	}

	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums[i+1:], temp, res)
		temp = temp[:len(temp)-1]
	}
}
