package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newhead := leetcode(head)

	return newhead
}

func leetcode(node *ListNode) *ListNode {
	if node.Next == nil {
		newhead := new(ListNode)
		newhead = node
		return newhead
	}

	last := leetcode(node.Next)
	last.Next = node
	node.Next = nil
	last = last.Next

	return last
}
