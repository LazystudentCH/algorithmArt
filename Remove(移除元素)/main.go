package Remove

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return k
}
