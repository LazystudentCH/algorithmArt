package main

import "fmt"

//面试官你好，我叫陈虎，目前大四，自学计算机大约一年半，有两份实习经验，两份实习都是后端开发，开发语言主要是Golang
func kmp(str1 string, str2 string) int {
	if len(str2) > len(str1) {
		return -1
	}
	arr := getLongestArr(str2)
	p1, p2 := 0, 0
	for p1 < len(str1) && p2 < len(str2) {
		if str1[p1] == str2[p2] {
			p1++
			p2++
		} else if arr[p2] == -1 {
			p1++
		} else {
			p2 = arr[p2]
		}
	}
	if p2 == len(str2) {
		return p1 - p2
	}
	return -1
}

//获取最长前缀后缀
func getLongestArr(str string) []int {
	if len(str) < 2 {
		return []int{0}
	}
	res := make([]int, len(str))
	res[0] = -1
	res[1] = 0
	i := 2
	cur := 0
	for i < len(str) {
		if str[i-1] == str[cur] {
			res[i] = cur + 1
			i++
		} else if cur > 0 {
			cur = res[cur]
		} else {
			res[i] = 0
			i++
		}
	}
	return res
}

func main() {
	a := "abcabcd"
	b := "abcd"
	fmt.Println(kmp(a, b))
}
