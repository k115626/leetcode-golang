package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 1
	for quick := 1; quick < len(nums); quick++ {
		if nums[quick] != nums[quick-1] {
			nums[slow] = nums[quick]
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
