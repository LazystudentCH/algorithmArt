package main

import (
	"container/heap"
	"fmt"
)

/*
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:
如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

*/

type minHeap []int
type maxHeap []int
type MedianFinder struct {
	minHeap minHeap
	maxHeap maxHeap
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := make(minHeap, 0)
	max := make(maxHeap, 0)
	heap.Init(&min)
	heap.Init(&max)
	return MedianFinder{
		minHeap: min,
		maxHeap: max,
	}
}

func (this *MedianFinder) AddNum(num int) {

	heap.Push(&this.maxHeap, num)
	heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))

	if len(this.maxHeap) < len(this.minHeap) {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}

	return
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) > len(this.minHeap) {
		return float64(this.maxHeap[0])
	}
	return (float64(this.maxHeap[0]) + float64(this.minHeap[0])) / 2
}

func main() {
	a := Constructor()
	a.AddNum(1)
	a.AddNum(2)
	fmt.Println(a.FindMedian())
	a.AddNum(3)
	fmt.Println(a.FindMedian())
}
