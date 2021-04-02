package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		currWidth := right - left
		currHeight := 0
		if height[left] > height[right] {
			currHeight = height[right]
			right--
		} else {
			currHeight = height[left]
			left++
		}
		temp := currHeight * currWidth
		if temp > res {
			res = temp
		}
	}
	return res
}

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(height)
	fmt.Println(res)
}
