package 不用加减乘除做加法

/*
写一个函数，求两个整数之和，要求在函数体内不得使用＋、－、×、÷ 四则运算符号。

样例
输入：num1 = 1 , num2 = 2

输出：3
*/

func add(num1 int, num2 int) int {
	for num2 != 0 {
		sum := num1 ^ num2
		temp := (num1 & num2) << 1
		num1 = sum
		num2 = temp
	}
	return num1
}
