package main

import "fmt"

func selectK(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	midIndex := partition(nums)

	if midIndex+1 == k {
		return nums[midIndex]
	}

	if midIndex+1 < k {
		return selectK(nums[midIndex+1:], k-midIndex-1)
	} else {
		return selectK(nums[:midIndex], k)
	}

}

func partition(nums []int) int {
	head, tail := 0, len(nums)-1
	x, i := nums[0], 1
	for head < tail {
		if x > nums[i] {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		}
	}
	return head
}

func main() {
	a := []int{3, 2, 1, 5, 6, 4}
	//partition(a)
	fmt.Println(selectK(a, 2))
	fmt.Println(a)
}
