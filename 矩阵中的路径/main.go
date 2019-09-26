package main

/*

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。

路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。

如果一条路径经过了矩阵中的某一个格子，则之后不能再次进入这个格子。

注意：

输入的路径不为空；
所有出现的字符均为大写英文字母；
样例
matrix=
[
  ["A","B","C","E"],
  ["S","F","C","S"],
  ["A","D","E","E"]
]

str="BCCE" , return "true"

str="ASAE" , return "false"

*/

func hasPath(matrix [][]byte, str string) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if dfs(matrix, str, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func dfs(matrix [][]byte, str string, length int, x, y int) bool {
	if matrix[x][y] != str[length] {
		return false
	}

	if length == len(str)-1 {
		return true
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	t := matrix[x][y]
	matrix[x][y] = '*'
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 0 && a < len(matrix) && b >= 0 && b < len(matrix[0]) {
			if dfs(matrix, str, length+1, a, b) {
				return true
			}
		}
	}
	matrix[x][y] = t
	return false
}

func main() {

}
