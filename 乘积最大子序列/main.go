package 乘积最大子序列

/*
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProduct(nums []int) int {
	Max := 1
	Min := 1
	m := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			Max, Min = Min, Max
		}
		Max = max(Max*nums[i], nums[i])
		Min = min(Min*nums[i], nums[i])
		m = max(m, Max)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
