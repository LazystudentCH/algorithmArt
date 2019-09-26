package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := deleteDuplicates(head)
	return h
}

func del(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	last := del(head.Next)
	if head.Val == last.Val {
		head.Next = last.Next
	}

	return head
}
