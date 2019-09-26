package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	p1 := 0
	p2 := len(str) - 1
	for p1 < p2 {
		if string(str[p1]) != string(str[p2]) {
			return false
		}
	}

	return true
}
