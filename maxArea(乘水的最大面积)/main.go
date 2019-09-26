package main

import "fmt"

func maxArea(height []int) int {
	r := len(height) - 1
	area := 0
	for i := 0; i < r; {
		area = Max((r-i)*Min(height[r], height[i]), area)
		if height[i] > height[r] {
			r--
		} else {
			i++
		}
	}

	return area
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}
