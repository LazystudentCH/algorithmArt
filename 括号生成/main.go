package main

func generateParenthesis(n int) []string {

}

func herper(left int, right int, n int, temp string, res *[]string) {
	if right == n {
		*res = append(*res, temp)
		return
	}

	if left < n {
		temp += "("
		herper(left+1, right, n, temp, res)
	}

	if right < left {
		temp += ")"
		herper(left, right+1, n, temp, res)
	}
}

//func generateParenthesis(n int) []string {
//	res := make([]string,0)
//	helper(0,0,n,"",&res)
//	return res
//}
//
//func helper(left,right,n int,temp string, res *[]string)  {
//	if right == n {
//		*res = append(*res,temp)
//		return
//	}
//
//	if left < n {
//		temp += "("
//		helper(left+1,right,n,temp,res)
//	}
//
//	if right < left {
//		temp += ")"
//		helper(left,right+1,n,temp,res)
//	}
//
//}
