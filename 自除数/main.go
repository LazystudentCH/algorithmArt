package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)

	for i := left; i <= right; i++ {
		nums := helper(i)
		flag := true
		for k := 0; k < len(nums); k++ {
			if i%nums[k] != 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}

func helper(num int) []int {
	if num < 10 {
		return []int{num}
	}
	res := make([]int, 0)

	Num := strconv.Itoa(num)
	n := []byte(Num)
	for i := 0; i < len(n); i++ {
		a, _ := strconv.Atoi(string(n[i]))
		res = append(res, a)
	}

	return res
}

func main() {
	fmt.Println()
}
