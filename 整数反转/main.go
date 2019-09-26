package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	ran := 2147483647
	str := strconv.Itoa(x)
	flag := false
	if strings.HasPrefix(str, "-") {
		flag = true
		str = strings.Replace(str, "-", "", -1)
	}

	str = reverseStr(str)
	res, _ := strconv.Atoi(str)
	if res >= ran {
		return 0
	} else {
		if flag {
			return -res
		}
		return res

	}

}

func reverseStr(str string) string {
	S := []byte(str)
	p1 := 0
	p2 := len(S) - 1

	for p1 < p2 {
		S[p1], S[p2] = S[p2], S[p1]
		p1++
		p2--
	}
	return string(S)
}

func main() {

	fmt.Println(reverse(1999999999))
}
