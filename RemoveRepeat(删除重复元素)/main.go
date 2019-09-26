package main

import "fmt"

// 0 0 1 1 -1 -1 2 3 3
func removeDuplicates(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func main() {
	b := []int{0, 0, 1, 1, 1, 1, 2, 3}
	a := removeDuplicates(b)
	fmt.Println(b, a)
}
