package main

import "fmt"

func parter(src string, parstr string) int {
	i, j := 0, 0
	p1 := 0
	S := []byte(src)
	Par := []byte(parstr)
	for p1 <= len(src)-1 {
		if j > len(parstr)-1 {
			return j - 1
		}

		if string(S[p1]) == string(Par[j]) {
			j++
			p1++
			continue
		} else {
			i++
			j = 0
			p1 = i
		}
	}
	return -1
}

func main() {
	fmt.Println(parter("abcabcabc", "abc"))
}
