package 打印两个链表的公共部分

import "fmt"

/*
链接：https://www.nowcoder.com/questionTerminal/8943eea40dbb4185b187d80fd050fee9
来源：牛客网

给定两个升序链表，打印两个升序链表的公共部分。

输入描述:
第一个链表的长度为 n。

第二个链表的长度为 m。

链表结点的值为 val。


输出描述:
输出一行整数表示两个升序链表的公共部分的值 (按升序输出)。
示例1
输入
4
1 2 3 4
5
1 2 3 5 6
输出
1 2 3
*/
//思路：两个指针，小的先走，如果相等，同时往后走

func Solution(l1, l2 []int) {
	p1, p2 := 0, 0
	for p1 < len(l1) && p2 < len(l2) {
		if l1[p1] < l2[p2] {
			p1++
		} else if l1[p1] > l2[p2] {
			p2++
		} else {
			fmt.Print(l1[p1])
			fmt.Print(" ")
			p1++
			p2++
		}
	}
}

func main() {
	n := 0
	fmt.Scan(&n)
	l1 := make([]int, 0)
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		l1 = append(l1, a)
	}
	m := 0
	fmt.Scan(&m)
	l2 := make([]int, 0)
	for i := 0; i < m; i++ {
		a := 0
		fmt.Scan(&a)
		l2 = append(l2, a)
	}
	Solution(l1, l2)
}
