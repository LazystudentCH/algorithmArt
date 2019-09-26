package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	if root.Val == sum {
		res = append(res, append([]int{}, root.Val))
	}

	l := pathSum(root.Left, sum-root.Val)
	for i := 0; i < len(l); i++ {
		res = append(res, append([]int{root.Val}, l[i]...))
	}
	r := pathSum(root.Right, sum-root.Val)
	for i := 0; i < len(r); i++ {
		res = append(res, append([]int{root.Val}, r[i]...))
	}

	return res
}

//思路 主要是遍历这一步,需要把之前存储的切片重新拷贝到结果切片中.
//如果需要在切片前加入元素可以将平时append两个参数对调
