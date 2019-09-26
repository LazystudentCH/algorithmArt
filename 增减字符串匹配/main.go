package main

func diStringMatch(S string) []int {
	i_c := 0
	d_c := 0
	I := 0
	D := 0
	s := []byte(S)
	s = append(s, s[len(s)-1])
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "D" {
			if d_c == 0 {
				d_c++
				res = append(res, len(s)-1)
				D = len(s) - 1
			} else {
				res = append(res, D-1)
				D = D - 1
			}
		} else {
			if i_c == 0 {
				i_c++
				res = append(res, I)
			} else {
				res = append(res, I+1)
				I++
			}
		}

	}

	return res
}
