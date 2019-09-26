package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	Max := 0
	l, r := 0, 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		tmp := make([]bool, len(s))
		dp[i] = tmp
	}
	for i := 0; i < len(dp); i++ {
		for j := i; j < len(dp); j++ {
			if i == j {
				dp[i][j] = true
			}
		}
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp) - 1; j > i; j-- {
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] == true {
				t := Max
				Max = max(Max, j-i+1)
				if t != Max {
					l = i
					r = j
				}
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return string(s[l : r+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindrome("aaaa"))
}
