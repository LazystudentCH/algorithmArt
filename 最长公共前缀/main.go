package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	com := ""
	if len(strs) == 0 {
		return com
	}

	com = strs[0]
	for i := 1; i < len(strs); i++ {
		p1 := 0
		p2 := 0
		for p1 < len(com) && p2 < len(strs[i]) {
			if com[p1] == strs[i][p2] {
				p1++
				p2++
			} else {
				break
			}
		}
		com = com[:p1]

	}

	return com
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
