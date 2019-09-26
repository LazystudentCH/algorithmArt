package main

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	A := 0
	B := 0
	m := make(map[string][]int, 0)
	m2 := make(map[string]int, 0)
	Sec := []byte(secret)
	for index, v := range Sec {
		m[string(v)] = append(m[string(v)], index)
	}

	Gue := []byte(guess)
	for index, v := range Gue {
		if len(m[string(v)]) != 0 {
			flag := false
			for _, k := range m[string(v)] {
				if k == index {
					flag = true
					A++
				}
			}

			if !flag {
				B++
			}
		}

		m2[string(v)]++
	}

	for k, v := range m2 {
		if v > len(m[k]) && len(m[k]) != 0 {

			B -= v - len(m[k])
			fmt.Println(B)
		}
	}

	a := strconv.Itoa(A)
	b := strconv.Itoa(B)
	return a + "A" + b + "B"
}

func main() {
	getHint("1123", "0111")
}
