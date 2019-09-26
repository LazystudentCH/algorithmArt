package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
	quickSort(nums)
	res := 999
	result := 0
	for i := 0; i < len(nums)-2; i++ {
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[head] + nums[tail] + nums[i]
			if sum == target {
				return sum
			}
			fmt.Println(head, tail, sum, res)
			if sum < target {
				head++
			} else {
				tail--
			}
			if min(res, abs(sum, target)) {
				res = abs(sum, target)
				result = sum
			}
		}
	}
	return result
}

func abs(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func min(a, b int) bool {
	if b < a {
		return true
	} else {
		return false
	}
}

func quickSort(nums []int) {
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

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}

func main() {
	fmt.Println(threeSumClosest([]int{-3, 0, 1, 2}, 1))
}
