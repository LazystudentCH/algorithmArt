package 构建乘积数组

/*
给定一个数组A[0, 1, …, n-1]，请构建一个数组B[0, 1, …, n-1]，其中B中的元素B[i]=A[0]×A[1]×… ×A[i-1]×A[i+1]×…×A[n-1]。

不能使用除法。

样例
输入：[1, 2, 3, 4, 5]

输出：[120, 60, 40, 30, 24]
思考题：

能不能只使用常数空间？（除了输出的数组之外）
*/
func multiply(A []int) []int {
	if len(A) == 0 {
		return nil
	}
	res := make([]int, 0)
	i := 0
	for len(res) != len(A) {
		temp := 1
		for k := 0; k < i; k++ {
			temp *= A[k]
		}
		for k := len(A) - 1; k > i; k-- {
			temp *= A[k]
		}
		i++
		res = append(res, temp)
	}
	return res
}
