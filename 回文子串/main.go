package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		helper(s, i, i, &res)
		helper(s, i, i+1, &res)
	}
	return res
}

func helper(s string, left, right int, res *int) {
	for left >= 0 && right < len(s) && string(s[left]) == string(s[right]) {
		*res++
		left--
		right++
	}
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
