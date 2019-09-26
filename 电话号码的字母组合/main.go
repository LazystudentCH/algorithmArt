package main

var M map[string][]string

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	M = make(map[string][]string, 0)
	M["2"] = []string{"a", "b", "c"}
	M["3"] = []string{"d", "e", "f"}
	M["4"] = []string{"g", "h", "i"}
	M["5"] = []string{"j", "k", "l"}
	M["6"] = []string{"m", "n", "o"}
	M["7"] = []string{"p", "q", "r", "s"}
	M["8"] = []string{"t", "u", "v"}
	M["9"] = []string{"w", "x", "y", "z"}

	a := []string{}
	for i := 0; i < len(digits); i++ {
		b := string(digits[i])
		a = append(a, b)
	}
	return helper(a)

}

func helper(str []string) []string {
	if str == nil {
		return nil
	}

	if len(str) == 1 {
		return M[str[0]]
	}

	res := []string{}
	l := helper(str[1:])
	for k := 0; k < len(M[str[0]]); k++ {
		for i := 0; i < len(l); i++ {
			a := M[str[0]][k]
			b := l[i]
			a += b
			res = append(res, a)
		}
	}

	return res
}
