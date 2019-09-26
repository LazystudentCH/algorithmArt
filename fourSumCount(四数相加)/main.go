package fourSumCount

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)

	for _, value := range C {
		for _, v := range D {
			m[v+value]++
		}
	}

	count := 0
	for _, v := range A {
		for _, value := range B {
			if m[0-(v+value)] != 0 {
				count += m[0-(v+value)]
			}
		}
	}

	return count
}
