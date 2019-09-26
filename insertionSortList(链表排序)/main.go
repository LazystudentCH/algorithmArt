package insertionSortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := sortList(head.Next)
	var pre *ListNode
	cur := node
	for {
		if head.Val < cur.Val {
			head.Next = cur
			if pre != nil {
				pre.Next = head
				node = pre
			} else {
				node = head
			}
			break
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return node

}
