package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	p1 := l1
	p2 := l2
	var temp int
	cur := dummy
	for p1 != nil || p2 != nil {
		if p1 != nil {
			temp += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			temp += p2.Val
			p2 = p2.Next
		}
		v := temp % 10
		node := &ListNode{v, nil}
		temp = temp / 10
		cur.Next = node
		cur = cur.Next
	}

	if temp == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next

}
