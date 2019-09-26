package main

import "fmt"

func quickSort(nums []int, l, r int) {
	if l < r {
		left, right := partition(nums, l, r)
		quickSort(nums, l, left-1)
		quickSort(nums, right+1, r)
	}
}

func partition(nums []int, l, r int) (int, int) {
	less := l - 1
	more := r
	for l < more {
		if nums[l] < nums[r] {
			nums[less+1], nums[l] = nums[l], nums[less+1]
			less++
			l++
		} else if nums[l] > nums[r] {
			nums[more-1], nums[l] = nums[l], nums[more-1]
			more--
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return less + 1, more
}

func main() {
	a := []int{6, 432, 432, 4, 2, 6, 2, 77, 2, 6, 10, 39, 49, 10, 10}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
