package main

func numJewelsInStones(J string, S string) int {
	m := make(map[string]int)
	j := []byte(J)
	s := []byte(S)
	res := 0
	for _, v := range j {
		m[string(v)] = 1
	}

	for i := 0; i < len(s); i++ {
		if m[string(s[i])] == 1 {
			res++
		}
	}

	return res
}
