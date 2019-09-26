package main

import "fmt"

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[cur] = i - cur
		}
		stack = append(stack, i)
	}

	fmt.Println(res)
	return res

}

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}
