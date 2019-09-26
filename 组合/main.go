package main

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

*/

var res [][]int
var limit int

func combine(n int, k int) [][]int {
	arr := make([]int, 0)
	for i := 1; i < n+1; i++ {
		arr = append(arr, i)
	}
	res = make([][]int, 0)
	limit = k
	helper(arr, 0, 0, []int{})
	return res
}

func helper(arr []int, k, i int, temp []int) {
	if len(temp) == limit || i == len(arr) {
		if len(temp) == limit {
			a := make([]int, len(temp))
			copy(a, temp)
			res = append(res, a)
		}
		return
	}

	helper(arr, k+1, i+1, append(temp, arr[i]))
	helper(arr, k, i+1, temp)
}

func main() {
	combine(4, 2)
}
