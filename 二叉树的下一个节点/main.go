package main

/*
给定一棵二叉树的其中一个节点，请找出中序遍历序列的下一个节点。

注意：

如果给定的节点是中序遍历序列的最后一个，则返回空节点;
二叉树一定不为空，且给定的节点一定不是空节点；
样例
假定二叉树是：[2, 1, 3, null, null, null, null]， 给出的是值等于2的节点。

则应返回值等于3的节点。

解释：该二叉树的结构如下，2的后继节点是3。
  2
 / \
1   3
*/

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}
	//有右子树
	if p.Right != nil {
		cur := p.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}

	//无右子树但是为父亲节点的左孩子
	if p.Father != nil && p.Father.Left == p {
		return p.Father
	}

	//无右子树但是为父亲节点的右孩子
	for p.Father != nil && p.Father.Right == p {
		p = p.Father
	}
	return p.Father

}

func main() {

}
