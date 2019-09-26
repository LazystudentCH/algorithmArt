package moris遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//四条规则:
//1.cur 无左子树，向右
//2.有左子树，找到左子树最有的节点
//1) 如果右子树为空，则右子树指向cur, cur = cur.left
//2) 右子树为cur, 右子树的右节点置为nil, cur = cur.right

//前序
func preorderTraversal(head *TreeNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	cur1 := head
	cur2 := cur1
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				res = append(res, cur1.Val)
				cur1 = cur1.Left
				continue
			} else {
				cur2.Right = nil
			}
		} else {
			res = append(res, cur1.Val)
		}
		cur1 = cur1.Right
	}
	return res
}

//中序
func inorderTraversal(head *TreeNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	cur1 := head
	cur2 := cur1
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				cur2.Right = nil
			}
		}
		res = append(res, cur1.Val)
		cur1 = cur1.Right
	}
	return res
}
