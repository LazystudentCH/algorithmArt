package main

import "fmt"

func deckRevealedIncreasing(deck []int) []int {
	sort(deck)
	res := make([]int, 0)
	//[2,13,3,11,5,17,7]
	//[2,3,5,7,11,13,17]
	for len(deck) != 1 {
		num := deck[len(deck)-1]
		res = append([]int{num}, res...)
		last := res[len(res)-1] //17
		res = res[:len(res)-1]
		res = append([]int{last}, res...)
		deck = deck[:len(deck)-1]
	}

	return append([]int{deck[0]}, res...)
}

func sort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	head, tail := 0, len(nums)-1
	x, i := nums[0], 1

	for head < tail {
		if x < nums[i] {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		}
	}

	sort(nums[:head])
	sort(nums[head+1:])
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}
