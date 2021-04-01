package main

// 以左下角的数为原点，定义一个指针指向该点
// 根据值的大小移动指针
// 当前值比数据小，则该列之上的元素都小，行等同
func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	width := len(matrix[0])
	currRow := height - 1
	currCol := 0
	for currRow >= 0 && currCol <= width-1 {
		currVal := matrix[currRow][currCol]
		switch {
		case currVal == target:
			return true
		case currVal < target:
			currCol += 1
		case currVal > target:
			currRow -= 1
		}
	}
	return false
}

func main() {

}
