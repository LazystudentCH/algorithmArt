package isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)
	S := []byte(s)
	T := []byte(t)

	for _, v := range S {
		m1[string(v)]++
	}

	for _, v := range T {
		m2[string(v)]++
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}

	}

	return true
}
