package main

func sortArrayByParity(A []int) []int {
	j, o := 0, len(A)-1

	for j < o {
		if A[j]%2 != 0 && A[o]%2 == 0 {
			A[j], A[o] = A[o], A[j]
			j++
			o--
		} else if A[j]%2 != 0 && A[o]%2 != 0 {
			o--
		} else if A[j]%2 == 0 && A[0]%2 == 0 {
			j++
		} else {
			j++
			o--
		}
	}

	return A
}
