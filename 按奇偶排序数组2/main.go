package main

func sortArrayByParityII(A []int) []int {
	j, o := 1, 0

	for o < len(A) && j < len(A) {
		if A[o]%2 != 0 && A[j]%2 == 0 {
			A[o], A[j] = A[j], A[o]
			o += 2
			j += 2
		} else if A[o]%2 == 0 && A[j]%2 == 0 {
			o += 2
		} else if A[o]%2 != 0 && A[j]%2 != 0 {
			j += 2
		} else {
			o += 2
			j += 2
		}
	}

	return A
}
