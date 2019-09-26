package 目标和

import "fmt"

/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组的长度不会超过20，并且数组中的值全为正数。
初始的数组的和不会超过1000。
保证返回的最终结果为32位整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, 2*(sum+1))
		dp = append(dp, tmp)
	}
	dp[0][sum] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j <= 2*(sum+1); j++ {
			if i < len(dp) && j <= 2*(sum+1) {
				dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(dp)-1][S]
}
