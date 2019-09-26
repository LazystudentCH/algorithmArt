package main

func removeOuterParentheses(S string) string {
	stack := make([]string, 0)
	s := []byte(S)
	result := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			if count != 0 {
				result += string(s[i])
			}
			count++
			stack = append(stack, string(s[i]))
		} else {
			if count != 1 {
				result += string(s[i])
			}
			stack = stack[:len(stack)-1]
			count--
		}
	}

	return result
}
