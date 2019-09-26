package main

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留。

样例1
输入：1->2->3->3->4->4->5

输出：1->2->5
样例2
输入：1->1->1->2->3

输出：2->3
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(head *ListNode) *ListNode {

}
