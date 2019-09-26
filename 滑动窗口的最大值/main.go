package 滑动窗口的最大值

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 {
		return nil
	}
	res := make([]int, 0)
	queue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		//如果双向队列的数字的最后一个数字小于当前数，需一直弹出
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		//检测头部是否过期
		if queue[0] == i-k {
			queue = queue[1:]
		}

		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
