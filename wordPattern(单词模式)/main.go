package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	Str := strings.Split(str, " ")
	m := make(map[string]string)
	pat := []byte(pattern)

	if len(pattern) == len(Str) {
		for index, v := range pat {
			if m[string(v)] == "" {
				m[string(v)] = Str[index]
			}
		}
	} else {
		return false
	}

	values := []string{}
	for _, v := range m {
		values = append(values, v)
	}

	if len(values) > 1 {
		if values[0] == values[1] {
			return false
		}
	}

	for index, v := range pat {
		if Str[index] != m[string(v)] {
			return false
		}
		fmt.Println(Str[index], m[string(v)])
	}

	return true
}

func main() {
	wordPattern("jequry", "jequry")
}
