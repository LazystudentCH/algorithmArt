package main

import (
	"container/heap"
	"fmt"
)

/*
题目： 一些项目要占用一个会议室宣讲，
会议室不能同时容纳两个项目的宣讲。给你每一个项目开始的时间和结束的时间(给你一个数组，里面是一个个具体的项目)，
你来安排宣讲的日程，要求会议室进行 的宣讲的场次最多。输出这个最多的宣讲场次。
*/

type meeting struct {
	start int
	end   int
}

type minHeap []meeting

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].end < m[j].end
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(meeting))
}

func (m *minHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x

}

func Solution(start []int, end []int) int {
	min := minHeap{}
	for i := 0; i < len(start); i++ {
		tmp := meeting{
			start: start[i],
			end:   end[i],
		}
		min = append(min, tmp)
	}
	heap.Init(&min)
	res := 1
	cur := min.Pop().(meeting).end
	for i := 0; i < len(min); i++ {
		m := min.Pop()
		if cur >= m.(meeting).end {
			res++
			cur = m.(meeting).end
		}
	}
	return res
}

/*
1   3
2   5
4   7
6   9
8   10
*/

func main() {
	fmt.Println(Solution([]int{1, 2, 4, 6, 8}, []int{3, 5, 7, 9, 10}))
}
