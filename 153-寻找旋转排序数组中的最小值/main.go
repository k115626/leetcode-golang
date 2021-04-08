package main

import "fmt"

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}

func main() {
	nums := []int{4, 5, 6, 7, 1, 2, 3}
	res := findMin(nums)
	fmt.Println(res)
}
