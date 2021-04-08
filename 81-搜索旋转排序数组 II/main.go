package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	start, mid, end := 0, 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[start] {
			start++
			continue
		}
		if nums[mid] > nums[start] {
			// 前半部分有序
			if nums[mid] > target && nums[start] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// 后半部分有序
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 3}
	target := 3
	res := search(nums, target)
	fmt.Println(res)
}
