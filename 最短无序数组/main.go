package main

func findUnsortedSubarray(nums []int) int {
	left, right := 0, -1
	max := nums[0]
	min := nums[len(nums)-1]
	r := len(nums) - 1
	for l := 0; l < len(nums); l++ {
		max = Max(max, nums[l])
		if max != nums[l] {
			right = l
		}

		min = Min(min, nums[r])
		if min != nums[r] {
			left = r
		}
		r--
	}

	return right - left + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
