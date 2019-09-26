package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dh := new(ListNode)
	dh.Next = head
	p := dh
	q := dh
	pre := dh

	for i := 0; i < n; i++ {
		q = q.Next
	}

	for {
		if q == nil {
			pre.Next = p.Next
			return dh.Next
		}

		pre = p
		p = p.Next
		q = q.Next
	}
}
