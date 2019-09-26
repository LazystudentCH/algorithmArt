package main

func toLowerCase(str string) string {
	s := []byte(str)

	for i := 0; i < len(s); i++ {
		if string(s[i]) >= "A" && string(s[i]) <= "Z" {

			s[i] += 32
		}
	}

	return string(s)
}
