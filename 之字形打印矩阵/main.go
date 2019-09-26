package main

import "fmt"

var flag = true

func PrintMatrix(m [][]int) {
	if len(m) == 0 || len(m[0]) == 0 {
		return
	}
	c1, r1 := 0, 0
	c2, r2 := 0, 0
	fmt.Print(m[r1][c1])
	fmt.Print(" ")
	for r2 < len(m) {
		if r1 != len(m)-1 {
			r1++
		} else {
			c1++
		}
		if c2 != len(m[0])-1 {
			c2++
		} else {
			r2++
		}
		PrintDJ(m, c1, r1, c2, r2)
		flag = !flag
	}
}

func PrintDJ(m [][]int, c1, r1, c2, r2 int) {
	if !flag {
		for c1 <= c2 {
			fmt.Print(m[r1][c1])
			fmt.Print(" ")
			r1--
			c1++
		}
	} else {
		for c1 <= c2 {
			fmt.Print(m[r2][c2])
			fmt.Print(" ")
			c2--
			r2++
		}
	}
}

func main() {
	r, c := 0, 0
	fmt.Scan(&r, &c)
	arr := make([][]int, 0)
	for i := 0; i < r; i++ {
		temp := make([]int, 0)
		for j := 0; j < c; j++ {
			a := 0
			fmt.Scan(&a)
			temp = append(temp, a)
		}
		arr = append(arr, temp)
	}
	PrintMatrix(arr)
}
