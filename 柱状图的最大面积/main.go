package main

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:

输入: [2,1,5,6,2,3]
输出: 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	area := 0
	minStack := make([]int, 0)
	minStack = append(minStack, -1)
	for i := 0; i < len(heights); i++ {
		for minStack[len(minStack)-1] != -1 && heights[minStack[len(minStack)-1]] >= heights[i] {
			left := minStack[len(minStack)-1]
			minStack = minStack[:len(minStack)-1]
			area = max(area, (i-minStack[len(minStack)-1]-1)*heights[left])
		}
		minStack = append(minStack, i)
	}

	for minStack[len(minStack)-1] != -1 {
		left := minStack[len(minStack)-1]
		minStack = minStack[:len(minStack)-1]
		area = max(area, heights[left]*(len(heights)-minStack[len(minStack)-1]-1))
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
