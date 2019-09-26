package main

import "fmt"

/*
给定无序数组arr， 已知arr中任意连个相邻的数都不相等。  写一个函数，只需返回arr中任意一个局部最小出现的位置

本题利用二分查找

   1 如果arr为空后者长度为0，返回-1

   2 如果arr长度为1或者arr[0]<arr[1]  返回0

  3 如果arr[N-1] < arr[N-2] 返回N-1

 4 如果长度大于2 且 arr的左右两头都不是局部最小，则令 left=1, right =N-2， 然后进行二分查找

   给定思路   mid=（left+right）/2

    如果 arr[mid] > arr[mid-1]  那么在 left......mid-1上肯定存在局部最小  所以令 right=mid-1

    如果 arr[mid] < arr[mid+1]  那么在 mid+1......right上肯定存在局部最小  所以令 left=mid+1

   上面两个都不满足，arr[mid] 就是局部最小  返回mid

    一直查找 知道left == right时停止  返回left就ok

由此可见！！！！！！ 二分查找并不是数组有序才能使用！！！！
*/
func getLessIndex(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	head := 0
	tail := len(arr) - 1
	if arr[0] < arr[1] {
		return 0
	}
	if arr[tail] < arr[tail-1] {
		return tail
	}
	for head < tail {
		mid := head + tail + 1>>1
		if arr[mid] > arr[mid-1] {
			tail = mid - 1
		} else {
			head = mid
		}

	}
	return head
}

func main() {
	a := []int{6, 4, 2, 3, 5, 9}
	fmt.Println(getLessIndex(a))
}
