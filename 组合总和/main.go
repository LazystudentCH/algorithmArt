package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	helper(candidates, []int{}, target, 0, &res)

	return res
}

func helper(nums, temp []int, target, start int, res *[][]int) {
	if target == 0 {
		a := make([]int, len(temp))
		copy(a, temp)
		*res = append(*res, a)
		return
	}

	if target < 0 || len(nums) == 0 {
		return
	}

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums, temp, target-nums[i], i, res)
		temp = temp[:len(temp)-1]
	}

}
func main() {

}
