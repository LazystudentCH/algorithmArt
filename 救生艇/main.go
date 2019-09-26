package main

func numRescueBoats(people []int, limit int) int {
	quickSort(people)
	res := 0
	l, r := 0, len(people)-1
	for l < r {
		if people[l]+people[r] <= limit {
			r--
			l++

		} else if people[l]+people[r] > limit {
			r--
		} else {
			l++
		}
		res++
	}

	return res
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	head, tail := 0, len(nums)-1
	x, i := nums[0], 1

	for head < tail {
		if nums[i] < x {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}

	quickSort(nums[0:head])
	quickSort(nums[head+1:])
}
