package main

import (
	"fmt"
	"sort"
)

// sort.Search(n, func)
// 返回0-n之间满足func==true的数的索引，采用二分查找
// 找到返回行索引，超出返回n, 小于返回0
// sort.SearchInts() 返回等同
// && 连接符有执行顺序，true的基础上向右判断

func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

func main() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	target := 12
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
