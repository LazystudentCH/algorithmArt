package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	data := helper(head)
	return Bst(data)
}

func Bst(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := new(TreeNode)
	mid := len(data) / 2
	root.Val = data[mid]
	root.Left = Bst(data[:mid])
	root.Right = Bst(data[mid+1:])
	return root

}

func helper(head *ListNode) []int {
	res := make([]int, 0)

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
