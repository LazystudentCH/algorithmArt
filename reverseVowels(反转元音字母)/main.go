package main

func reverseVowels(s string) string {
	m := make(map[string]int, 5)
	m["a"] = 1
	m["e"] = 1
	m["i"] = 1
	m["o"] = 1
	m["u"] = 1
	m["A"] = 1
	m["E"] = 1
	m["I"] = 1
	m["O"] = 1
	m["U"] = 1

	r := len(s) - 1
	data := []byte(s)
	for i := 0; i < len(data); {
		if i >= r {
			break
		}

		if m[string(data[i])] == 0 {
			i++
			continue
		}

		if m[string(data[r])] == 0 {
			r--
			continue
		}

		data[i], data[r] = s[r], s[i]
		i++
		r--
	}

	return string(data)
}

func main() {

}
