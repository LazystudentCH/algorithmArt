package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	count := len(lists)

	for {
		miner := 999
		changeIndex := 0
		nil_count := 0
		for j := 0; j < len(lists); j++ {
			if lists[j] == nil {
				nil_count++
				continue
			}
			if min(lists[j].Val, miner) {
				miner = lists[j].Val
				changeIndex = j
			}
		}
		if count == nil_count {
			break
		}
		if lists[changeIndex] != nil {
			lists[changeIndex] = lists[changeIndex].Next
		}
		node := &ListNode{
			Val: miner,
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func min(a, b int) bool {
	if a <= b {
		return true
	} else {
		return false
	}
}
func main() {

}
