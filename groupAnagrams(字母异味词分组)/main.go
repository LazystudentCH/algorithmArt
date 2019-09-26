package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	A := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	for index, v := range A {
		m[v] = a[index]
	}

	m1 := make(map[int][]string)
	for _, v := range strs {
		temp := []byte(v)
		j := cal(temp, m)
		m1[j] = append(m1[j], v)
	}

	result := [][]string{}
	for _, v := range m1 {
		result = append(result, v)
	}

	return result

}

func cal(data []byte, d map[string]int) int {
	res := 1
	for _, v := range data {
		res *= d[string(v)]
	}

	return res
}
