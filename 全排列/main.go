package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	helper(nums, []int{}, &res)
	return res
}

func helper(nums []int, temp []int, res *[][]int) {
	if len(nums) == 0 {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
		//fmt.Println(*res)
		return
	}

	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp = append(temp, nums[i])
		copy(a, nums)
		//fmt.Println(nums)
		helper(append(a[:i], a[i+1:]...), temp, res)
		temp = temp[:len(temp)-1]
	}

}

func main() {
	fmt.Println(permute([]int{5, 4, 2, 6}))

}
