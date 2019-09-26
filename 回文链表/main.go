package 回文链表

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//思路：快慢指针，当快指针走到最后的时候，慢指针在中间，反转后半部分，然后进行对比
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, low := head, head
	count := 0
	for fast != nil {
		fast = fast.Next
		count++
	}
	if count == 1 {
		return true
	}
	if count == 2 {
		return head.Val == head.Next.Val
	}

	fast = head
	if count%2 == 0 {
		//偶数
		for fast.Next != nil {
			fast = fast.Next
			low = low.Next
			if fast.Next != nil {
				fast = fast.Next
			} else {
				break
			}
		}
		pre, cur := low, low.Next
		next := cur.Next
		for next != nil {
			cur.Next = pre
			pre = cur
			cur = next
			next = next.Next
		}
		cur.Next = pre
		//从头开始
		for head.Next != fast {
			if head.Val != fast.Val {
				return false
			}
			head = head.Next
			fast = fast.Next
		}
		return head.Val == head.Next.Val
	} else {
		//奇数
		for fast.Next != nil {
			fast = fast.Next
			low = low.Next
			if fast.Next != nil {
				fast = fast.Next
			} else {
				break
			}
		}
		pre, cur := low, low.Next
		next := cur.Next
		for next != nil {
			cur.Next = pre
			pre = cur
			cur = next
			next = next.Next
		}
		cur.Next = pre
		//从头开始
		for head != fast {
			if head.Val != fast.Val {
				return false
			}
			head = head.Next
			fast = fast.Next
		}
	}
	return true

}
