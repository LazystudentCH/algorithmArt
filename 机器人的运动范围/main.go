package main

import (
	"fmt"
)

/*
地上有一个 m 行和 n 列的方格，横纵坐标范围分别是 0∼m−1 和 0∼n−1。

一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格。

但是不能进入行坐标和列坐标的数位之和大于 k 的格子。

请问该机器人能够达到多少个格子？

样例1
输入：k=7, m=4, n=5

输出：20
样例2
输入：k=18, m=40, n=40

输出：1484

解释：当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。
      但是，它不能进入方格（35,38），因为3+5+3+8 = 19。
注意:

0<=m<=50
0<=n<=50
0<=k<=100
*/

//深度优先
/*
func movingCount(threshold, rows, cols int) int {
	if rows == 0 || cols == 0 {
		return 0
	}
	maps := make([][]bool,0)
	for i := 0 ; i < rows; i++{
		temp := make([]bool,cols)
		maps = append(maps,temp)
	}
	res := 0
	dfs(maps,&res,0,0,threshold)
	return res
}

func dfs(maps [][]bool,res *int,x,y,threshold int)  {
	if helper(x,y) > threshold || maps[x][y] == true{
		return
	}

	*res++
	maps[x][y] = true
	dx := []int{-1,0,1,0}
	dy := []int{0,1,0,-1}
	for i := 0 ; i < 4 ; i++{
		a := x+dx[i]
		b := y+dy[i]
		if a >=0 && a< len(maps) && b >= 0 && b < len(maps[0]) && maps[a][b] != true{
			dfs(maps,res,a,b,threshold)
		}
	}
}

func helper(x,y int) int {
	x1 := (x/10 + (x%10))
	y1 := (y/10 + (y%10))
	return x1+y1
}
*/
//宽度优先
type point struct {
	x int
	y int
}

func movingCount(threshold, rows, cols int) int {
	if rows == 0 || cols == 0 {
		return 0
	}
	maps := make([][]bool, 0)
	for i := 0; i < rows; i++ {
		temp := make([]bool, cols)
		maps = append(maps, temp)
	}
	res := 0
	queue := make([]point, 0)
	queue = append(queue, point{x: 0, y: 0})
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if helper(p.x, p.y) > threshold || maps[p.x][p.y] == true {
			continue
		} else {
			maps[p.x][p.y] = true
			res++
		}
		for i := 0; i < 4; i++ {
			p2 := point{x: p.x + dx[i], y: p.y + dy[i]}
			if p2.x >= 0 && p2.x < len(maps) && p2.y >= 0 && p2.y < len(maps[0]) {
				queue = append(queue, p2)
			}
		}
	}
	return res
}

func helper(x, y int) int {
	x1 := (x/10 + (x % 10))
	y1 := (y/10 + (y % 10))
	return x1 + y1
}

func main() {
	fmt.Println(movingCount(3, 13, 14))
}
