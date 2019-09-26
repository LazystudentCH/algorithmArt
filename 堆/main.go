package main

import "fmt"

func makeHeap(nums []int) {
	for i := (len(nums) / 2) - 1; i >= 0; i-- {
		MinHeapShiftDown(nums, i, len(nums)-1)
	}

}

func MinHeapShiftDown(nums []int, index int, length int) {
	l_child := 2*index + 1
	r_child := 2*index + 2
	if l_child >= length {
		return
	}

	//寻找最小值
	min := l_child
	if r_child > length {
		min = l_child
	} else {
		if nums[l_child] > nums[r_child] {
			min = r_child
		}
	}

	if nums[index] <= nums[min] {
		return
	}

	nums[index], nums[min] = nums[min], nums[index]
	MinHeapShiftDown(nums, min, length)

}

func sort(nums []int) {
	makeHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		MinHeapShiftDown(nums, 0, i-1)
	}
}

func main() {
	a := []int{3, 4, 6, 1, 77, 2, 10}
	sort(a)
	fmt.Println(a)
}
