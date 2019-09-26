package main

func sortedSquares(A []int) []int {
	res := make([]int, 0)
	k := len(A) - 1

	for i := 0; i < k; {
		if A[i]*A[i] < A[k]*A[k] {
			res = append([]int{A[k] * A[k]}, res...)
			k--
		} else {
			res = append([]int{A[i] * A[i]}, res...)
			i++
		}
	}

	return res
}
