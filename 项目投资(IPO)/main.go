package main

import (
	"container/heap"
	"fmt"
)

/*
假设 力扣（LeetCode）即将开始其 IPO。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。

给定若干个项目。对于每个项目 i，它都有一个纯利润 Pi，并且需要最小的资本 Ci 来启动相应的项目。最初，你有 W 资本。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择最多 k 个不同项目的列表，以最大化最终资本，并输出最终可获得的最多资本。

示例 1:

输入: k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].

输出: 4

解释:
由于你的初始资本为 0，你尽可以从 0 号项目开始。
在完成后，你将获得 1 的利润，你的总资本将变为 1。
此时你可以选择开始 1 号或 2 号项目。
由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。


注意:

假设所有输入数字都是非负整数。
表示利润和资本的数组的长度不超过 50000。
答案保证在 32 位有符号整数范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ipo
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Process struct {
	C int
	V int
}

type minHeap []Process

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].C < h[j].C
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	(*h) = append((*h), x.(Process))
}

func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

type maxHeap []Process

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].V > h[j].V
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	(*h) = append((*h), x.(Process))
}

func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	c := minHeap{}
	v := maxHeap{}
	for i := 0; i < len(Profits); i++ {
		p := Process{
			C: Capital[i],
			V: Profits[i],
		}
		c = append(c, p)
	}
	heap.Init(&c)
	heap.Init(&v)
	for k > 0 {
		for len(c) > 0 && W >= c[0].C {
			heap.Push(&v, heap.Pop(&c))
		}
		if len(v) == 0 {
			break
		}
		W += heap.Pop(&v).(Process).V
		k--
	}

	return W
}

func main() {
	fmt.Println(findMaximizedCapital(1, 0, []int{1, 2, 3}, []int{0, 1, 2}))
}
