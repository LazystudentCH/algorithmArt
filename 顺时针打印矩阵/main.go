package main

func printMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row1 := 0
	col1 := 0
	row2 := len(matrix) - 1
	col2 := len(matrix[0]) - 1
	res := make([]int, 0)

	for col1 <= col2 && row1 <= row2 {
		temp := helper(matrix, row1, col1, row2, col2)
		col1++
		col2--
		row1++
		row2--
		res = append(res, temp...)
	}

	return res
}

func helper(matrix [][]int, row1, col1, row2, col2 int) []int {
	res := make([]int, 0)
	if row1 == row2 {
		for i := col1; i <= col2; i++ {
			res = append(res, matrix[row1][i])
		}
		return res
	}
	if col1 == col2 {
		for i := row1; i <= row2; i++ {
			res = append(res, matrix[i][col1])
		}
		return res
	}

	r, c := row1, col1
	for c != col2 {
		res = append(res, matrix[row1][c])
		c++
	}

	for r != row2 {
		res = append(res, matrix[r][c])
		r++
	}

	for c != col1 {
		res = append(res, matrix[row2][c])
		c--
	}

	for r != row1 {
		res = append(res, matrix[r][col1])
		r--
	}
	return res

}

func main() {

}
