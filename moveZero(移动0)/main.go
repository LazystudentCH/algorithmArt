package main

func moveZeroes(nums []int) {

	u := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[u], nums[i] = nums[i], nums[u]
			u++
		}
	}

}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
