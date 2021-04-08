package main

import (
	"fmt"
)

// n个元素的数组，缺失的最小的正数肯定在 1-n+1之间
// 元素可以重复
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < 1 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num < n+1 {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for k, v := range nums {
		if v > 0 {
			return k + 1
		}
	}
	return n + 1
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	nums := []int{3, 4, -1, 1}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
