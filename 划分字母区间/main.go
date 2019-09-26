package main

import "fmt"

func partitionLabels(S string) []int {
	s := []byte(S)
	m := make(map[string]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		m[string(s[i])] = i
	}
	fmt.Println(m)
	start := 0
	end := m[string(s[start])]
	for i := 0; i < len(s); {
		for k := i; k <= end; k++ {
			if m[string(s[k])] > end {
				end = m[string(s[k])]
			}
		}
		fmt.Println(end, i)
		res = append(res, end-i+1)
		i = end + 1
		if i >= len(s) {
			break
		}
		end = m[string(s[i])]

	}

	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
