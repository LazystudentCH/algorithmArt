package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	helper(nums, []int{}, k, n, &res)

	return res
}

func helper(nums, temp []int, k, target int, res *[][]int) {
	if len(temp) == k && target == 0 {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
		return
	}

	if len(nums) == 0 && target < 0 {
		return

	}

	for i := 0; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums[i+1:], temp, k, target-nums[i], res)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
