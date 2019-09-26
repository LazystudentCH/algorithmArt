package 二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			res = append(res, stack[len(stack)-1].Val)
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
