package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = helper([]byte(words[i]))
	}

	return strings.Join(words, " ")
}

func helper(word []byte) string {
	head, tail := 0, len(word)-1
	for head < tail {
		word[head], word[tail] = word[tail], word[head]
		head++
		tail--
	}

	return string(word)
}
