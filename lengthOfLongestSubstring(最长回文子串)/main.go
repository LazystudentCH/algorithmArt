package main

import "strings"

type huiwen struct {
	length int
	start  int
	end    int
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	S := []byte(s)
	b := []string{""}
	for i := 0; i < len(S); i++ {
		b = append(b, string(S[i]))
	}
	b = append(b, "")
	Str := []byte(strings.Join(b, "#"))
	// Str 长度为2len(s)+1
	hw := new(huiwen)
	for i := 0; i < len(Str); i++ {
		a := i
		hw = Max(Helper(Str, a), hw)
	}

	if hw.length == 0 || hw.length == 1 {
		return string(S[0])
	}
	return strings.Replace(string(Str[hw.start:hw.end]), "#", "", -1)

}

func Max(a, b *huiwen) *huiwen {
	if a.length > b.length {
		return a
	}

	return b
}

func Helper(data []byte, i int) *huiwen {
	length := 0
	a := new(huiwen)
	for k := 0; k <= i && k < len(data)-i; k++ {
		if string(data[i-k]) == string(data[i+k]) {
			length++
			a.start = i - k
			a.end = i + k + 1
		} else {
			a.start = i - k + 2
			a.end = i + k - 1
			break
		}
	}
	a.length = length - 1

	return a
}
