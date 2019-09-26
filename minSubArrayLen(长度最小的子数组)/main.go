package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	j := -1
	sum := 0
	length := -1
	for i := 0; i < len(nums); {

		if sum >= s {
			if length == -1 {
				length = len(nums[i : j+1])
				continue
			}
			length = Min(length, len(nums[i:j+1]))
			sum -= nums[i]
			i++
		} else {
			if j < len(nums)-1 {
				j++
				sum += nums[j]
				continue
			}

			i++

		}
	}

	if length == -1 {
		length = 0
	}
	return length
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(res)
	//
	//a:=[]int{1,2,3,4,5}
	//fmt.Println(len(a[2:4]))
}
