package main

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	head, tail := 0, len(data)-1
	x, i := data[0], 1

	for head < tail {
		if data[i] > x {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}

	QuickSort(data[:head])
	QuickSort(data[head+1:])

}

func main() {

}
