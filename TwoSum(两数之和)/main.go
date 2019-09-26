package main

// 时间复杂度O(n)
// 空间复杂度O(n)
// 利用map
//func twoSum(numbers []int, target int) []int {
//	m := make(map[int]string,0)
//
//	for index,value := range numbers{
//			if m[value] == "" {
//				m[value] = strconv.Itoa(index)
//			}else{
//				if value*2 == target{
//					index1 ,_:= strconv.Atoi(m[value])
//					return []int{index1+1,index+1}
//				}
//			}
//	}
//
//
//
//	for k,v := range m {
//		if m[target-k] != "" && target/2 != k{
//			s , _:= strconv.Atoi(v)
//			index ,_:= strconv.Atoi(m[target-k])
//			if index > s {
//				return []int{s+1,index+1}
//			}else{
//				return []int{index+1,s+1}
//			}
//		}
//	}
//
//	return nil
//}
//

//时间复杂度O(nlogn)
//空间复杂度O(1)
//二分搜索,如果碰到有序数组,优先考虑二分搜索
//func twoSum(numbers []int, target int) []int {
//	for i := 0; i<len(numbers); i ++ {
//		as := target - numbers[i]
//		l := i+1
//		r := len(numbers)-1
//		res := Search(numbers,l,r,as)
//		if res != -1 {
//			return []int{i+1,res+1}
//		}
//
//	}
//
//	return nil
//}
//
//func Search(nums []int,l int , r int, as int) int {
//	for ;l <= r ; {
//		mid := l + (r-l)/2
//		if nums[mid] == as {
//			return mid
//		}
//
//		if as > nums[mid] {
//			l = mid +1
//		}
//
//		if as < nums[mid]{
//			r = mid -1
//		}
//	}
//
//	return -1
//}

//指针碰撞
func twoSum(numbers []int, target int) []int {
	r := 0
	for i := 0; i < len(numbers); {
		if numbers[i]+numbers[r] == target {
			return []int{i + 1, r + 1}
		}

		if numbers[i]+numbers[r] < target {
			i++
			continue
		}

		if numbers[i]+numbers[r] > target {
			r++
		}
	}

	return nil
}

func main() {

}
