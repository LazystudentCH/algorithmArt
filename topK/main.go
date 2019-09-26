package main

func topK(nums []int, k int) []int {
	res := make([]int, 0)
	temp := nums[:k]
	makeHeap(temp)
	for i := k; i < len(nums); i++ {
		if nums[i] > temp[0] {
			res = append(res, temp[0])
			temp[0] = nums[i]
			minHeapShiftDown(temp, 0, len(temp)-1)
		}
	}

	return res
}

func makeHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		minHeapShiftDown(nums, i, len(nums)-1)
	}
}

func minHeapShiftDown(nums []int, index int, length int) {
	if index >= length {
		return
	}
	l_child := 2*index + 1
	r_child := 2*index + 2
	min := l_child

	if l_child > length {
		//叶子节点
		return
	}

	if r_child > length {
		min = l_child
	} else {
		if nums[l_child] > nums[r_child] {
			min = r_child
		}
	}

	nums[index], nums[min] = nums[min], nums[index]
	minHeapShiftDown(nums, min, length)
}
