package main

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

*/
func minPathSum(arr [][]int) int {
	dp := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		temp := make([]int, len(arr[0]))
		dp = append(dp, temp)
	}
	dp[0][0] = arr[0][0]
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = arr[0][i] + dp[0][i-1]
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = arr[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + arr[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {

}
