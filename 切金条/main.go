package main

import (
	"container/heap"
	"fmt"
)

/*
题目说明：
一块金条切成两半，是需要花费和长度数值一样的铜板的。比如
长度为20的 金条，不管切成长度多大的两半，都要花费20个铜
板。一群人想整分整块金 条，怎么分最省铜板？
例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为
10+20+30=60. 金条要分成10,20,30三个部分。 如果， 先把长
度60的金条分成10和50，花费60 再把长度50的金条分成20和30，花费50 一共花费110铜板。
但是如果， 先把长度60的金条分成30和30，花费60 再把长度30
金条分成10和20，花费30 一共花费90铜板。
输入一个数组，返回分割的最小代价。

算法思想：
基于哈夫曼编码的思想，
(1) 首先构造小根堆
(2) 每次取最小的两个数（小根堆），使其代价最小。并将其和加入到小根堆中
(3) 重复(2)过程，直到最后堆中只剩下一个节点
————————————————
版权声明：本文为CSDN博主「imprincess」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/imprincess/article/details/82055111
*/
type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append((*h), x.(int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func lessMoney(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	a := Heap{}
	for i := 0; i < len(nums); i++ {
		a = append(a, nums[i])
	}
	heap.Init(&a)
	for len(a) > 1 {
		nums1 := heap.Pop(&a)
		nums2 := heap.Pop(&a)
		x := nums1.(int) + nums2.(int)
		res += x
		heap.Push(&a, x)
	}

	return res
}

func main() {
	fmt.Println(lessMoney([]int{20, 30}))
}
