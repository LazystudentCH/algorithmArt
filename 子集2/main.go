package main

func subsetsWithDup(nums []int) [][]int {
	sort(nums)
	fmt.Println(nums)
	res := make([][]int, 0)
	helper(nums, []int{}, 0, &res)
	return append(res, []int{})
}

func helper(nums []int, temp []int, start int, res *[][]int) {
	if len(temp) != 0 {
		a := make([]int, len(temp))
		copy(a, temp)
		*res = append(*res, a)
	}

	if len(nums) == 0 {
		return
	}

	for i := start; i < len(nums); i++ {
		if i != start && (nums[i] == nums[i-1]) {
			continue
		}
		temp = append(temp, nums[i])
		helper(nums, temp, i+1, res)
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
