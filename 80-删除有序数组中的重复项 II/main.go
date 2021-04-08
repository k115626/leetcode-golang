package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	slow, quick := 2, 2
	for quick < len(nums) {
		if nums[slow-2] != nums[quick] {
			nums[slow] = nums[quick]
			slow++
		}
		quick++
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
