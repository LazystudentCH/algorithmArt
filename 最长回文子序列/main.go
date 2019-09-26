package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	S := []byte(s)
	res := ""
	for i := 1; i < len(S); i++ {
		a := helper(S, i, i)
		b := helper(S, i, i+1)
		res = max(max(a, b), res)
	}
	return res
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	} else {
		return b
	}
}

func helper(Str []byte, left, right int) string {
	for left >= 0 && right < len(Str) {
		if left-1 < 0 || right+1 > len(Str)-1 {
			return string(Str[left : right+1])
		}
		if string(Str[left-1]) == string(Str[right+1]) {
			left--
			right++
		} else {
			break
		}
	}

	return string(Str[left : right+1])
}

func main() {
	fmt.Println(longestPalindrome("ac"))
}
