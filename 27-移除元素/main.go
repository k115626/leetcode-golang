package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0
	for quick := 0; quick < len(nums); quick++ {
		if nums[quick] != val {
			nums[slow] = nums[quick]
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res)
}
