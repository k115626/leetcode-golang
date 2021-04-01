package main

import "fmt"

// 1 2 3 4 5
func climbStairs(n int) int {
	// dp 存储每个台阶的可能的数
	dp := make([]int, n)
	for i := 1; i < n+1; i++ {
		switch i {
		case 1:
			dp[i-1] = 1
		case 2:
			dp[i-1] = 2
		default:
			dp[i-1] = dp[i-2] + dp[i-3]
		}
	}
	return dp[len(dp)-1]
}

func main() {
	n := 5
	res := climbStairs(n)
	fmt.Println(res)
}
