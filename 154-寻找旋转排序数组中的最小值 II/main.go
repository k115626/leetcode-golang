package main

import "fmt"

/*
start mid end
1. mid > end ==> min 在后半部
2. mid = end ==> 无法确定
3. mid < end ==> min 前半部
*/

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}
	return nums[start]
}

func main() {
	nums := []int{3, 3, 1, 3}
	res := findMin(nums)
	fmt.Println(res)
}
