package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	m1 := make(map[string]string, 0)
	p := []byte(pattern)

	for i := 0; i < len(words); i++ {
		w := []byte(words[i])
		if len(w) != len(pattern) {
			continue
		}

		arr := make(map[string]bool, 0)
		for i := 0; i < len(w); i++ {
			arr[string(w[i])] = false
		}

		for i := 0; i < len(p); i++ {
			m1[string(p[i])] = ""
		}

		flag := true
		for k := 0; k < len(w); k++ {
			if m1[string(p[k])] == "" && arr[string(w[k])] == false {
				m1[string(p[k])] = string(w[k])
				arr[string(w[k])] = true
				continue
			}
			fmt.Println(m1)
			if m1[string(p[k])] != string(w[k]) {
				flag = false
			}
		}

		if flag {
			res = append(res, words[i])
		}
	}

	return res
}

func main() {
	findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
}
