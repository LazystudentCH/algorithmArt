package main

func judgeCircle(moves string) bool {
	m1 := make(map[string]int)
	s := []byte(moves)
	for i := 0; i < len(moves); i++ {
		m1[string(s[i])]++
	}

	if m1["U"] == m1["D"] && m1["L"] == m1["R"] {
		return true
	} else {
		return false
	}
}
