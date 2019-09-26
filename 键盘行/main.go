package main

func findWords(words []string) []string {
	if words == nil {
		return nil
	}
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		w := []byte(words[i])
		if isLine(w) {
			result = append(result, words[i])
		}
	}

	return result
}

func isLine(word []byte) bool {
	m1 := []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	m2 := []string{"A", "S", "D", "F", "G", "H", "J", "K", "L", "a", "s", "d", "f", "g", "h", "j", "k", "l"}
	m3 := []string{"Z", "X", "C", "V", "B", "N", "M", "z", "x", "c", "v", "b", "n", "m"}
	count := 0
	for i := 0; i < len(word); i++ {
		for k := 0; k < len(m1); k++ {
			if string(word[i]) == m1[k] {
				count++
			}
		}
	}
	if count == len(word) {
		return true
	}
	count = 0

	for i := 0; i < len(word); i++ {
		for k := 0; k < len(m2); k++ {
			if string(word[i]) == m2[k] {
				count++
			}
		}
	}
	if count == len(word) {
		return true
	}
	count = 0

	for i := 0; i < len(word); i++ {
		for k := 0; k < len(m3); k++ {
			if string(word[i]) == m3[k] {
				count++
			}
		}
	}
	if count == len(word) {
		return true
	}

	return false
}
