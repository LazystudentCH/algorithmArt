package main

func helper(s string, temp string, res *[]string) {
	if temp != "" {
		*res = append(*res, temp)
	}

	for i := 0; i < len(s); i++ {
		temp += string(s[i])
		helper(string(s[i+1:]), temp, res)
		temp = temp[:len(temp)-1]
	}
}
