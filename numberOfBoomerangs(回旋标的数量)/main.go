package main

import (
	"fmt"
	"math"
)

func numberOfBoomerangs(points [][]int) int {

	res := 0
	for index, p := range points {
		m := make(map[float64]int)
		points[index], points[0] = points[0], points[index]
		for i := 1; i < len(points); i++ {
			d := cal(p, points[i])
			m[d]++
		}

		for _, v := range m {
			if v > 1 {
				temp := 1
				for i := v; i > v-2; i-- {
					temp *= i
				}
				res += temp
			}
		}
	}

	return res
}

func cal(a, b []int) float64 {

	return math.Sqrt(math.Pow(float64(b[1]-a[1]), 2) + math.Pow(float64(b[0]-a[0]), 2))
}

func main() {
	res := numberOfBoomerangs([][]int{[]int{0, 0}, []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}})
	fmt.Println(res)
}
