package 岛数量

/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++
				helper(grid, i, j)
			}
		}
	}
	return res
}

func helper(arr [][]byte, i, j int) {
	if i < 0 || i >= len(arr) || j < 0 || j >= len(arr[0]) || arr[i][j] != '1' {
		return
	}
	arr[i][j] = '2'
	helper(arr, i, j+1)
	helper(arr, i, j-1)
	helper(arr, i+1, j)
	helper(arr, i-1, j)
}
