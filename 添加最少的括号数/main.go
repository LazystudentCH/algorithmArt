package main

func minAddToMakeValid(S string) int {
	stack := make([]string, 0)
	count := 0
	s := []byte(S)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				count++
			}
		}
	}

	return count + len(stack)
}
