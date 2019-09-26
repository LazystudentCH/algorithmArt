package main

import "fmt"

func stoneGame(piles []int) bool {
	l, r := 0, len(piles)
	t, k := 0, 0
	flag := true
	for l < r {

		if flag {
			flag = false
			if piles[l] > piles[r] {
				t += piles[l]
				l++
			} else {
				t += piles[r]
				r--
			}

		} else {
			flag = true
			if piles[l] > piles[r] {
				k += piles[l]
				l++
			} else {
				k += piles[r]
				r--
			}
		}
	}

	return t > k
}
