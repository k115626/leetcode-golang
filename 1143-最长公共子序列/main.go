package main

import "fmt"

// 动态规划
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			switch {
			case text1[i-1] == text2[j-1]:
				dp[i][j] = dp[i-1][j-1] + 1
			default:
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	text1 := "abc"
	text2 := "def"
	res := longestCommonSubsequence(text1, text2)
	fmt.Println(res)
}
