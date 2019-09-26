package main

func uniqueMorseRepresentations(words []string) int {
	m1 := make(map[string]string, 0)
	m1 = map[string]string{"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".", "f": "..-.", "g": "--.", "h": "....", "i": "..", "j": ".---", "k": "-.-", "l": ".-..", "m": "--", "n": "-.", "o": "---", "p": ".--.", "q": "--.-", "r": ".-.", "s": "...", "t": "-", "u": "..-", "v": "...-", "w": ".--", "x": "-..-", "y": "-.--", "z": "--.."}
	m2 := make(map[string]int, 0)
	for _, v := range words {
		res := ""
		for j := 0; j < len([]byte(v)); j++ {
			res += m1[string([]byte(v)[j])]
		}
		m2[res]++
	}

	count := 0
	for _, _ = range m2 {
		count++
	}
	return count
}
