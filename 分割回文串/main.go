package main

import "fmt"

func partition(s string) [][]string {
	res := [][]string{}
	helper(s, []string{}, &res)
	return res
}

func helper(s string, temp []string, res *[][]string) {
	if len(s) == 0 {
		//fmt.Println(temp)
		b := make([]string, len(temp))
		copy(b, temp)
		*res = append(*res, b)
		fmt.Println(res)
		return
	}

	S := []byte(s)
	for i := 1; i < len(S)+1; i++ {
		if IsHw(S[:i]) {
			temp = append(temp, string(S[:i]))
			helper(string(S[i:]), temp, res)
			temp = temp[:len(temp)-1]
		}

	}

}

func IsHw(s []byte) bool {
	r := len(s) - 1

	for i := 0; i < r; i++ {
		if s[i] != s[r] {
			return false
		}
		r--
	}

	return true
}

func main() {
	fmt.Println(partition("abbbcc"))
}
