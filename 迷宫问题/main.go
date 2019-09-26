package main

import "fmt"

type point struct {
	x int
	y int
}

var maps [][]int = [][]int{{0, 0, 0, 0, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 1, 0, 0},
	{0, 1, 1, 0, 1},
	{0, 0, 0, 0, 0}}

func bfs(arr [][]int) {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return
	}
	n := len(arr)
	m := len(arr[0])
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	tmp := make([][]bool, n)
	for i := 0; i < n; i++ {
		a := make([]bool, m)
		tmp[i] = a
	}
	queue := make([]point, 0)
	queue = append(queue, point{0, 0})
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			tmp[queue[i].x][queue[i].y] = true
			fmt.Println("(", queue[i].x, ",", queue[i].y, ")")
			for j := 0; j < 4; j++ {
				a := queue[i].x + dx[j]
				b := queue[i].y + dy[j]
				if a >= 0 && a < n && b >= 0 && b < m && arr[a][b] != 1 && tmp[a][b] == false {
					queue = append(queue, point{a, b})
				}
			}
		}
		queue = queue[length:]
	}
}

func main() {
	bfs(maps)
}
