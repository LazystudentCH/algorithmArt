package 二叉树的后序遍历序列

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。

如果是则返回true，否则返回false。

假设输入的数组的任意两个数字都互不相同。

样例
输入：[4, 8, 6, 12, 16, 14, 10]

输出：true
*/
//思路：
func verifySequenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return true
	}
	root := sequence[len(sequence)-1]
	index := 0
	flag := 0
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] > root && flag == 0 {
			index = i
			flag = 1
		} else if sequence[i] < root && flag == 1 {
			return false
		}
	}
	return verifySequenceOfBST(sequence[:index]) && verifySequenceOfBST(sequence[index:len(sequence)-1])
}
