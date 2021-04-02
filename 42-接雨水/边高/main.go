package main

import "fmt"

// 将问题展开来看
// i处可以接到水的值是其两侧最大高度的最小值和本身的差
// 定义两个数组
// leftmax 记录i处的左侧最大值
// rightmax 记录i处的右侧最大值
// i 处的雨水 = min(leftmax[i], rightmax[i]) - height[i]
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	res := 0
	for i := 0; i < len(height); i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
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
