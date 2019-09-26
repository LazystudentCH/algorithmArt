package main

import (
	"strings"
)

/*
请实现一个函数，把字符串中的每个空格替换成"%20"。

你可以假定输入字符串的长度最大是1000。
注意输出字符串的长度可能大于1000。

样例
输入："We are happy."

输出："We%20are%20happy."
*/

func replaceSpaces(str string) string {
	return strings.Replace(str, " ", "%20", -1)
}

func main() {

}
