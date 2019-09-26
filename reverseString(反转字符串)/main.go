package main

import "fmt"

func reverseString(s []byte) {
	r := len(s) - 1

	for i := 0; i < len(s); i++ {
		if i == r {
			break
		}

		if r-i == 1 {
			s[i], s[r] = s[r], s[i]
			break
		}

		s[i], s[r] = s[r], s[i]
		r--
	}
}

func main() {
	s := []byte("abcde")
	fmt.Println(string(s[0]))
}
