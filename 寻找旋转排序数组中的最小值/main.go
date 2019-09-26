package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] < nums[r] {
		return nums[0]
	}
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid
		}
	}
	return nums[r]

}

func main() {

}
