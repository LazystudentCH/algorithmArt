package main

func smallestRangeI(A []int, K int) int {
	sort(A)
	res := A[len(A)-1] - K - (A[0] + K)
	if res < 0 {
		return 0
	}

	return res
}

func sort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	head, tail := 0, len(nums)-1
	x, i := nums[0], 1

	for head < tail {
		if x < nums[i] {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		}
	}

	sort(nums[:head])
	sort(nums[head+1:])
}
