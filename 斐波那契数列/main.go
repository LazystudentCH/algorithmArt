package main

/*

输入一个整数 n ，求斐波那契数列的第 n 项。

假定从0开始，第0项为0。(n<=39)

样例
输入整数 n=5

返回 5
*/

func Fibonacci(n int) int {
	n1 := 1
	n2 := 1
	res := 0
	for i := 2; i < n; i++ {
		res = n1 + n2
		n1 = n2
		n2 = res
	}
	return res
}

func main() {

}
