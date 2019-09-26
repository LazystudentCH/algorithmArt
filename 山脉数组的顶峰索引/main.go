package main

func peakIndexInMountainArray(A []int) int {
	index := 0
	for i := 0; i < len(A); i++ {
		if A[i] < A[i+1] {
			index++
		} else {
			return index
		}
	}

	return index
}
