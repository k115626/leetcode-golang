package main

import "fmt"

// 边界
// y==m-1 只能向右移动
// x==n-1 只能向下移动
// 可以向下/向右移动
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// finish 坐标定义为(1,1)
	// start 坐标定义为(m, n)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			switch {
			case x == 0 && y == 0:
				dp[x][y] = 1
			case x == 0 && y == 1:
				dp[x][y] = 1
			case x == 1 && y == 0:
				dp[x][y] = 1
			case x == 0:
				dp[x][y] = dp[x][y-1]
			case y == 0:
				dp[x][y] = dp[x-1][y]
			default:
				dp[x][y] = dp[x-1][y] + dp[x][y-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	m, n := 3, 3
	res := uniquePaths(m, n)
	fmt.Println(res)
}
