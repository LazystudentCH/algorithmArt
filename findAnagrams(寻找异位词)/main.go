package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var M1 = make(map[string]int)
	res := []int{}
	data := []byte(s)
	P := []byte(p)
	for _, v := range P {
		M1[string(v)]++
	}

	l := 0
	r := len(P) - 1

	for l < len(data) {
		if isAnagrams(data[l:r+1], M1) {
			res = append(res, l)
		}
		l++
		r++

		if r == len(data) {
			break
		}
	}

	return res
}

func isAnagrams(s []byte, d map[string]int) bool {
	m2 := make(map[string]int)
	for _, v := range s {
		if d[string(v)] == 0 {
			return false
		} else {
			m2[string(v)]++
		}
	}

	for k, _ := range m2 {
		if m2[k] == d[k] {
			continue
		} else {
			return false
		}
	}

	return true

}

func main() {
	res := findAnagrams("cbaebabacd", "abc")
	fmt.Println(res)
}
