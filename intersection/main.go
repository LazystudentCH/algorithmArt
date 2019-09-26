package intersection

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}

	m2 := make(map[int]int)
	for _, v := range nums2 {
		if m2[v] == 0 {
			m2[v] = 1
		}
	}

	m1 := make(map[int]int)
	for _, v := range nums1 {
		if m1[v] == 0 {
			m1[v] = 1
		}

		if m2[v] == 1 {
			res = append(res, v)
			m2[v]++
		}
	}

	return res
}
