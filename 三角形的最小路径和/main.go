package 三角形的最小路径和

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		tmp := make([]int, len(triangle[i]))
		dp[i] = tmp
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		if i == 1 {
			dp[i][0] = dp[0][0] + triangle[i][0]
			dp[i][1] = dp[0][0] + triangle[i][1]
			continue
		}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	res := dp[len(dp)-1][0]
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		res = min(res, dp[len(dp)-1][i])
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
