package main

import "fmt"

/*
最大值减去最小值小于或等于 k 的子数组数量
思路：两个双向队列，一个最大，一个最小
遍历数组，从某一个下标开始,右边界一直往外扩，扩到不能阔的时候,res += x+1  x为右边界
然后start++
*/

func test(arr []int, k int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if isValid(arr[i:j+1], k) {
				res++
			}
		}
	}
	return res
}

func isValid(nums []int, k int) bool {
	Max := nums[0]
	Min := nums[0]
	for i := 0; i < len(nums); i++ {
		Max = max(Max, nums[i])
		Min = min(Min, nums[i])
	}
	return Max-Min <= k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func subArray(nums []int, k int) int {
	res := 0
	max := make([]int, 0)
	min := make([]int, 0)
	l, r := 0, 0
	for l < len(nums) {
		for r < len(nums) {
			for len(max) != 0 && nums[max[len(max)-1]] <= nums[r] {
				max = max[:len(max)-1]
			}
			max = append(max, r)
			for len(min) != 0 && nums[min[len(min)-1]] >= nums[r] {
				min = min[:len(min)-1]
			}
			min = append(min, r)
			if nums[max[0]]-nums[min[0]] > k {
				break
			}
			r++
		}
		if max[0] == l {
			max = max[1:]
		}
		if min[0] == l {
			min = min[1:]
		}
		res += r - l
		l++
	}
	return res
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	fmt.Println(test(a, 3))
	fmt.Println(a)
	fmt.Println(subArray(a, 3))
}
