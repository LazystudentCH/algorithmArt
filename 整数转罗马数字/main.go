package main

var Count map[int]string

func init() {
	Count = make(map[int]string, 0)
	Count[1] = "I"
	Count[4] = "IV"
	Count[5] = "V"
	Count[9] = "IX"
	Count[10] = "X"
	Count[40] = "XL"
	Count[50] = "L"
	Count[90] = "XC"
	Count[100] = "C"
	Count[400] = "CD"
	Count[500] = "D"
	Count[900] = "CM"
	Count[1000] = "M"
}

func intToRoman(num int) string {
	if Count[num] != "" {
		return Count[num]
	}

	res := ""
	for k, v := range Count {
		a := num / k
		num -= a * k
		intToRoman(num)
		for i := 0; i < a; i++ {
			res += v
		}
	}
	return res
}

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/integer-to-roman
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	intToRoman()
}
