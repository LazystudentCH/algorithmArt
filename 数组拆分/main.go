package main

func arrayPairSum(nums []int) int {
	sort(nums)
	head, tail := 0, len(nums)-1
	res := 0
	for head < tail {
		if tail-head <= 1 {
			res += min(nums[tail], nums[tail-1])
			break
		}
		res += min(nums[head], nums[head+1]) + min(nums[tail], nums[tail-1])
		head += 2
		tail -= 2
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
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
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}

	sort(nums[:head])
	sort(nums[head+1:])
}
