package main

func repeatedNTimes(A []int) int {
	m1 := make(map[int]int, 0)
	for i := 0; i < len(A); i++ {
		m1[A[i]]++
	}

	for k, v := range m1 {
		if v == len(A)/2 {
			return k
		}
	}

	return 0
}
