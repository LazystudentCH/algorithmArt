package partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var h1 = new(ListNode)
	var h2 = new(ListNode)
	var node1 = h1
	var node2 = h2

	for head.Next != nil {
		if head.Val < x {
			node1.Next = head
			head = head.Next
			node1.Next = nil
		} else {
			node2.Next = head
			head = head.Next
			node2.Next = nil
		}
	}

	node1.Next = h2.Next

	return h1.Next
}
