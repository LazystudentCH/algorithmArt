package main

import "fmt"

// 3,2,1,5,6,4    k = 2
// 2 1 3 5 6 4
// 4 5 6
func findKthLargest(nums []int, k int) int {
	for {
		if len(nums) == 1 {
			return nums[0]
		}

		index := QuickSort(nums)

		if index == len(nums)-1 {
			if k == 1 {
				return nums[index]
			} else {
				nums = nums[:index]
				k--
				continue
			}
		}
		fmt.Println(nums)

		if len(nums[index+1:]) == k-1 {
			return nums[index]
		}

		if len(nums[index+1:]) >= k {
			nums = nums[index+1:]
			continue
		}

		if len(nums[index+1:]) < k {
			k = k - len(nums[index+1:]) - 1
			nums = nums[:index]
		}
	}

}

func QuickSort(nums []int) int {
	s := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[s] {
			nums[i], nums[s] = nums[s], nums[i]
			nums[s+1], nums[i] = nums[i], nums[s+1]
			s++
		}
	}

	return s
}

//2 4 1 3 0       2    3
func main() {

	nums := []int{5, 2, 4, 1, 3, 6, 0}
	//s := QuickSort(nums)
	//fmt.Println(nums,s)
	findKthLargest(nums, 4)

}
