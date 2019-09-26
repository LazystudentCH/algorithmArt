package 二叉树的路径和

/*
输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。

从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。

样例
给出二叉树如下所示，并给出num=22。
      5
     / \
    4   6
   /   / \
  12  13  6
 /  \    / \
9    1  5   1

输出：[[5,4,12,1],[5,6,6,5]]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func findPath(root *TreeNode, sum int) [][]int {
	res = make([][]int, 0)
	dfs(root, sum, []int{})
	return res

}

func dfs(root *TreeNode, sum int, temp []int) {
	if root == nil {
		return
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		temp = append(temp, root.Val)
		a := make([]int, len(temp))
		copy(a, temp)
		res = append(res, a)
		return
	}
	temp = append(temp, root.Val)
	dfs(root.Left, sum-root.Val, temp)
	dfs(root.Right, sum-root.Val, temp)
	return
}
