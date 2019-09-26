package 分隔链表

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//思路:三个指针 small mid big  分别代表小于大于等于，然后依此便利链表，放在对应的区域，然后再连起来
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	small_head := new(ListNode)
	small := small_head
	big_head := new(ListNode)
	big := big_head
	for head != nil {
		if head.Val > x {
			big.Next = head
			big = big.Next
		}
		if head.Val < x {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}

	small.Next = big_head.Next
	return small_head.Next
}
