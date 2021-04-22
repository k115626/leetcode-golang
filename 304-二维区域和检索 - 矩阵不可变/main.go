package main

// 一维前缀和
/*
是将每一行按之前的一维一样求前缀和，
后面将多行的和拿到再求和，时间复杂度是O(m)
*/
// 二维前缀和
/*
直接按二维的来，时间复杂度是O(1)
*/
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	height, width := len(matrix), len(matrix[0])
	sums := make([][]int, height+1)
	for i := 0; i < len(sums); i++ {
		sums[i] = make([]int, width+1)
	}
	for m, n := range matrix {
		for x, y := range n {
			sums[m+1][x+1] = sums[m+1][x] + sums[m][x+1] + y - sums[m][x]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {

}
