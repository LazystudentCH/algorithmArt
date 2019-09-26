package 股票的最大利润

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖 一次 该股票可能获得的利润是多少？

例如一只股票在某些时间节点的价格为[9, 11, 8, 5, 7, 12, 16, 14]。

如果我们能在价格为5的时候买入并在价格为16时卖出，则能收获最大的利润11。

样例
输入：[9, 11, 8, 5, 7, 12, 16, 14]

输出：11

由于只允许做一次股票买卖交易，枚举每一天作为卖出的日子，买入日子一定在卖出日子之前，为了获利最多，
希望买入的日子的股票价格尽可能低。用minnum[i]记录第0-i天的最低价格，则在第i天卖出的最大获利为nums[i]-minnum[i]，枚举i找到最大获利
*/
func maxDiff(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		if dp[i-1] < nums[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(nums[i]-dp[i], res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
