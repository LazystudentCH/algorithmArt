package main

/*
输入一个链表的头结点，按照 从尾到头 的顺序返回节点的值。

返回的结果用数组存储。

样例
输入：[2, 3, 5]
返回：[5, 3, 2]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListReversingly(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append([]int{head.Val}, res...)
		head = head.Next
	}
	return res
}
