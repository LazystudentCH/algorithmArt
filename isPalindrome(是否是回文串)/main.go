package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	for index, value := range s {
		if !(("a" <= string(value) && string(value) <= "z") || ("A" <= string(value) && string(value) <= "Z")) && !("0" <= string(value) && string(value) <= "9") {
			s = strings.Replace(s, string(s[index]), " ", 1)
		}
	}

	r := len(s) - 1
	for i := 0; i < len(s); {

		if string(s[i]) == " " {
			i++
			continue
		}

		if string(s[r]) == " " {
			r--
			continue
		}

		if strings.ToUpper(string(s[i])) != strings.ToUpper(string(s[r])) {
			return false
		}

		if r == i || r-i == 1 && (r != len(s)-1 && i != 0) {
			return true
		}

		i++
		r--
	}

	return true

}

func main() {

	s := "`l;`` 1o1 ??;l`" //a 97  z 122 A 65 Z90
	isPalindrome(s)
}
