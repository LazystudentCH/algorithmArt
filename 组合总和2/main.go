package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort(candidates)
	helper(candidates, []int{}, target, 0, &res)
	return res
}

func helper(nums, temp []int, target, start int, res *[][]int) {
	if target == 0 {
		a := make([]int, len(temp))
		copy(a, temp)
		*res = append(*res, a)
		return
	}

	if len(nums) == 0 || target < 0 {
		return
	}

	for i := start; i < len(nums); i++ {
		if i != start && (nums[i] == nums[i-1]) {
			continue
		}
		temp = append(temp, nums[i])
		helper(nums, temp, target-nums[i], i+1, res)
		temp = temp[:len(temp)-1]

	}
}

func sort(nums []int) {
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
			nums[head], nums[i] = nums[i], nums[head]
			i++
			head++
		}
	}

	sort(nums[:head])
	sort(nums[head+1:])
}

func main() {
	a := []int{1, 2, 2}
	sort(a)
	fmt.Println(a)
}
