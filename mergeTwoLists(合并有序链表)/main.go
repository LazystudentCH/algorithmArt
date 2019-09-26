package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	res := merge(l1, l2)
	return res

}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var dh *ListNode
	var h *ListNode
	dh = new(ListNode)
	h = dh.Next

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			h = l2
			l2 = l2.Next
			h.Next = nil
			h = h.Next
		} else {
			h = l1
			l1 = l1.Next
			h.Next = nil
			h = h.Next
		}
	}

	if l1 == nil {
		h.Next = l2
	} else {
		h.Next = l1
	}

	return dh.Next
}
