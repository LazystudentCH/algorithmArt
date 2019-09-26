package main

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	m1 := make(map[string]int, 0)
	l1 := make([][]string, 0)
	for _, v := range s {
		m1[string(v)]++
	}

	for k, v := range m1 {
		temp := []string{}
		for i := 0; i < v; i++ {
			temp = append(temp, k)
		}

		l1 = append(l1, temp)
	}

	fmt.Println(l1)
	for i := 0; i < len(l1); i++ {
		for k := 0; k < len(l1)-i-1; k++ {
			if len(l1[k+1]) > len(l1[k]) {
				l1[k], l1[k+1] = l1[k+1], l1[k]
			}
		}
	}

	fmt.Println(l1)
	result := ""
	for i := 0; i < len(l1); i++ {
		a := i
		result += strings.Join(l1[a], "")
	}

	return result
}

func main() {
	fmt.Println(frequencySort("raaeaedere"))
}
