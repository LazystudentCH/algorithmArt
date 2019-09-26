package main

func sortColors(nums []int) {
	zero := -1
	two := len(nums)

	for i := 0; i < two; {
		if nums[i] == 0 {
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
			continue
		}

		if nums[i] == 1 {
			i++
			continue
		}

		if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
			continue

		}

	}
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
