package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	new_nums := make([]int, 0)

	k := 0
	s := 0

	if len(nums1) == 0 {
		nums1 = nums2
		return
	}

	if len(nums2) == 0 {
		return
	}

	for i := 0; i < m; {
		if k == n {
			s = i
			break
		}

		if nums1[i] < nums2[k] {
			new_nums = append(new_nums, nums1[i])
			i++
		}

		if nums1[i] >= nums2[k] {
			new_nums = append(new_nums, nums2[k])
			k++
		}

		s = i
	}

	if s == m && k < n {
		for i := k; i < n; i++ {
			new_nums = append(new_nums, nums2[i])
		}
	}

	if k == n && s < m {
		for i := s; i < m; i++ {
			new_nums = append(new_nums, nums1[i])
		}
	}

	for i := 0; i < len(new_nums); i++ {
		nums1[i] = new_nums[i]
	}

}
func main() {
	merge([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3)
}
