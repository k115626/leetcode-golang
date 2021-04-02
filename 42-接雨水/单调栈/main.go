package main

import "fmt"

func trap(height []int) int {
	stack := make([]int, 0)
	res := 0
	for k, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			popNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			currWidth := k - stack[len(stack)-1] - 1
			currHeight := min(height[stack[len(stack)-1]], v) - height[popNum]
			res += currHeight * currWidth
		}
		stack = append(stack, k)
	}
	return res
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2}
	res := trap(height)
	fmt.Println(res)
}
