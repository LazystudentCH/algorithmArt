package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int][]*TreeNode

func init() {
	m = make(map[int][]*TreeNode)
}

func allPossibleFBT(N int) []*TreeNode {
	if m[N] != nil {
		return m[N]
	}

	if N == 1 {
		root := &TreeNode{
			Val: 0,
		}
		return []*TreeNode{root}
	}

	if N == 2 {
		return nil
	}

	l_size := 1
	r_size := N - l_size - 1
	res := make([]*TreeNode, 0)
	for r_size >= 1 {
		l := allPossibleFBT(l_size)
		for i := 0; i < len(l); i++ {
			r := allPossibleFBT(r_size)
			for k := 0; k < len(r); k++ {
				root := &TreeNode{
					Val: 0,
				}
				root.Left = l[i]
				root.Right = r[k]
				res = append(res, root)
			}
		}

		l_size += 2
		r_size = N - l_size - 1
		m[N] = res
	}

	return res
}
