package main

import "container/heap"

type element struct {
	pl  int
	val int
}

type elementheap []element

func (h elementheap) Len() int {
	return len(h)
}

func (h elementheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *elementheap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

func (h elementheap) Less(i, j int) bool {
	return h[i].pl >= h[j].pl
}

func (h *elementheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	que := &elementheap{}
	heap.Init(que)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for k, v := range m {
		e := element{
			pl:  v,
			val: k,
		}
		if len(*que) != k {
			heap.Push(que, e)
		} else {
			x := heap.Pop(que)
			if x.(element).pl < v {
				heap.Push(que, e)
			} else {
				heap.Push(que, x)
			}
		}
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(que).(element).val)
	}
	return result
}
