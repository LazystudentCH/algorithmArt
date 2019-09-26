package main

func MoreThanHalf(nums []int) int {
	return partition(nums)
}

func partition(nums []int) int {
	head, tail := 0, len(nums)-1
	x, i := nums[0], 1
	for head < tail {
		if x > nums[i] {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}

	return tail
}

func MoreThanHalf2(nums []int) int {
	num := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
			count++
			continue
		}

		if num != nums[i] {
			count--
		} else {
			count++
		}
	}

	return num
}

func main() {

}
