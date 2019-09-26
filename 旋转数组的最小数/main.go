package main

/*

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。

输入一个升序的数组的一个旋转，输出旋转数组的最小元素。

例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。

数组可能包含重复项。

注意：数组内所含元素非负，若数组大小为0，请返回-1。

样例
输入：nums=[2,2,2,0,1]

输出：0
*/

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}

	return -1
}

func main() {

}
