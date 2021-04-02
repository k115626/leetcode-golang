package main

import "fmt"

// 总体积减去柱子的体积就是水的容量
func trap(height []int) int {
	leftBorder, rightBorder := 0, len(height)-1
	sumVolume, barVolume := 0, 0
	for _, v := range height {
		barVolume += v
	}
	floorNum := 1
	for leftBorder <= rightBorder {
		for leftBorder <= rightBorder && height[leftBorder] < floorNum {
			leftBorder++
		}
		for leftBorder <= rightBorder && height[rightBorder] < floorNum {
			rightBorder--
		}
		sumVolume += rightBorder - leftBorder + 1
		floorNum++
	}
	waterVolume := sumVolume - barVolume
	return waterVolume
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	res := trap(height)
	fmt.Println(res)
}
