package 求1_2___n

/*
求1+2+…+n,要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

样例
输入：10

输出：55
最直接的想法就是用递归，sum(n) = n+sum(n-1)，但是要注意终止条件，由于求的是1+2+…+n的和，所以需要在n=0的时候跳出递归，但是题目要求不能使用if,while等分支判断，可以考虑利用&&短路运算来终止判断。

时间复杂度分析：递归，复杂度为O(n)O(n)。
*/

func getSum(n int) int {
	var sum int
	ff(n, &sum)
	return sum
}

func ff(n int, result *int) int {
	*result += n
	_ = n > 0 && ff(n-1, result) > 0
	return n
}
