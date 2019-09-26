package main

/*
给你一根长度为 n 绳子，请把绳子剪成 m 段（m、n 都是整数，2≤n≤58 并且 m≥2）。

每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]k[1] … k[m] 可能的最大乘积是多少？

例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到最大的乘积18。

样例
输入：8

输出：18
*/
func maxProductAfterCutting(length int) int {
	if length <= 3 {
		return length - 1
	}
	res := 1
	if length%3 == 1 {
		res = 4
		length = length - 4
	} else if length%3 == 2 {
		res = 2
		length = length - 2
	}

	for length > 0 {
		res *= 3
		length -= 3
	}
	return res
}

func main() {

}
