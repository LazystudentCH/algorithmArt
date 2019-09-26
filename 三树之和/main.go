package main

import "fmt"

func threeSum(nums []int) [][]int {
	quickSort(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high, sum := i+1, len(nums)-1, 0-nums[i]
		for low < high {
			if sum == (nums[high] + nums[low]) {
				res = append(res, []int{nums[i], nums[high], nums[low]})
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else {
				if sum+nums[low]+nums[high] < 0 {
					low++
				} else {
					high--
				}
			}
		}
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
		if x < nums[i] {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			i++
			head++
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
