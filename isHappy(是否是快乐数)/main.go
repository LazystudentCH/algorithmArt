package main

import (
	"fmt"
	"math"
	"strconv"
)

func isHappy(n int) bool {
	return cal(n, 0)
}

func cal(n int, count int) bool {
	if n == 1 {
		return true
	}

	if count >= 50 {
		return false
	}

	N := Conv(n)
	sum := 0
	for _, v := range N {
		temp, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(temp), 2))
	}

	//fmt.Println(sum)
	count++
	return cal(sum, count)
}

func Conv(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func main() {
	res := isHappy(19)
	fmt.Println(res)
}
