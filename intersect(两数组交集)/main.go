package intersect

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}

	m2 := make(map[int]int)
	for _, v := range nums2 {
		m2[v]++
	}

	m1 := make(map[int]int)
	for _, v := range nums1 {
		m1[v]++

	}

	for _, v := range nums1 {
		if m2[v] != 0 {
			pl := Min(m1[v], m2[v])
			for i := 0; i < pl; i++ {
				res = append(res, v)
			}
		}
	}

	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
