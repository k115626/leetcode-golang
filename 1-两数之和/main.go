package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if _, ok := m[other]; ok {
			return []int{m[other], i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{5, 2, 3}
	target := 5
	res := twoSum(nums, target)
	fmt.Println(res)
}
