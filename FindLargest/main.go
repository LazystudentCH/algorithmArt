package main

type MaxHeap struct {
	data []int
}

func (self *MaxHeap) Add(data int) {
	self.data = append(self.data, data)
	self.add(len(self.data) - 1)
}

func (self *MaxHeap) ParentIndex(index int) int {
	return (index - 1) / 2
}

func (self *MaxHeap) LeftIndex(index int) int {
	return index*2 + 1
}

func (self *MaxHeap) RightIndex(index int) int {
	return index*2 + 2
}

func (self *MaxHeap) add(index int) {
	if index == 0 {
		return
	}

	if self.data[self.ParentIndex(index)] > self.data[index] {
		return
	}

	self.data[index], self.data[self.ParentIndex(index)] = self.data[self.ParentIndex(index)], self.data[index]
	self.add(self.ParentIndex(self.ParentIndex(index)))

}

func (self *MaxHeap) GetMaxNum() int {
	root := self.data[0]
	self.data[0] = self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	self.shiftdown(0)
	return root
}

func (self *MaxHeap) shiftdown(index int) {
	for {
		if self.LeftIndex(index) >= len(self.data) {
			break
		}

		if self.RightIndex(index) >= len(self.data) {
			if self.data[self.LeftIndex(index)] > self.data[index] {
				self.data[index], self.data[self.LeftIndex(index)] = self.data[self.LeftIndex(index)], self.data[index]
			}

			break
		}

		max_child := max(self.data[self.LeftIndex(index)], self.data[self.RightIndex(index)])
		if max_child == 0 {
			if self.data[index] < self.data[self.LeftIndex(index)] {
				self.data[index], self.data[self.LeftIndex(index)] = self.data[self.LeftIndex(index)], self.data[index]
				index = self.LeftIndex(index)
			} else {
				break
			}
		} else {
			if self.data[index] < self.data[self.RightIndex(index)] {
				self.data[index], self.data[self.RightIndex(index)] = self.data[self.RightIndex(index)], self.data[index]
				index = self.RightIndex(index)
			} else {
				break
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return 0
	}
	return 1
}

func Heapify(data []int) *MaxHeap {
	res := &MaxHeap{data: data}
	parent := res.ParentIndex(len(data) - 1)
	for i := parent; i >= 0; i-- {
		a := i
		res.shiftdown(a)
	}

	return res
}

func findKthLargest(nums []int, k int) int {
	heap := Heapify(nums)
	res := 0
	for i := 0; i < k; i++ {
		res = heap.GetMaxNum()
	}

	return res
}
