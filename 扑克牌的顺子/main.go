package 扑克牌的顺子

/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。

2～10为数字本身，A为1，J为11，Q为12，K为13，大小王可以看做任意数字。

为了方便，大小王均以0来表示，并且假设这副牌中大小王均有两张。

样例1
输入：[8,9,10,11,12]

输出：true
样例2
输入：[0,8,9,11,12]

输出：true
*/
func isContinuous(nums []int) bool {
	if len(nums) != 5 {
		return false
	}
	zeroCount := 0
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		} else {
			if nums[i] == nums[i+1] {
				//出现对子
				return false
			}
			if nums[i+1]-nums[i]-1 > zeroCount {
				//差大于0的个数
				return false
			}
		}
		zeroCount -= nums[i+1] - nums[i] - 1
	}
	return true
}

func quickSort(nums []int, l, r int) {
	if l < r {
		left, right := partition(nums, l, r)
		quickSort(nums, l, left-1)
		quickSort(nums, right+1, r)
	}
}

func partition(nums []int, l, r int) (int, int) {
	less := l - 1
	more := r
	for l < more {
		if nums[l] < nums[r] {
			nums[less+1], nums[l] = nums[l], nums[less+1]
			l++
			less++
		} else if nums[l] > nums[r] {
			nums[more-1], nums[l] = nums[l], nums[more-1]
			more--
		} else {
			l++
		}
	}
	nums[r], nums[more] = nums[more], nums[r]
	return less + 1, more
}
