package main

import "fmt"

func Reverse() {

}

func Mergesort(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}
	mid := len(nums) / 2
	left, b := Mergesort(nums[:mid])
	right, c := Mergesort(nums[mid:])
	res, a := Merge(left, right)
	return res, a + b + c

}

func Merge(left, right []int) ([]int, int) {
	res := make([]int, 0)
	count := 0
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			res = append(res, left[l])
			l++
		} else {
			count += len(left[l:])
			res = append(res, right[r])
			r++
		}
	}

	res = append(res, left[l:]...)
	res = append(res, right[r:]...)

	return res, count
}

func main() {
	a := []int{1, 5, 6, 11, 3, 55, 2}
	fmt.Println(Mergesort(a))
}
