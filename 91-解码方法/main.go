package main

import "fmt"

/*
动态规划：
dp[i] =
	1. 添加的是一位可解码，则dp[i] = dp[i-1]
	2. 添加的与倒数第二位组成两位可解码，则dp[i] = dp[i-2]
*/
// 为什么dp[i-1] 和 dp[i-2] 没有重复的方案？
// -- 两个方案中只要有一个元素不同则就是两种不同的方案
func numDecodings(s string) int {
	// dp = [1, 0, 0, 0]
	//          |
	//          V
	// 实际是从第二个位置开始，第一个位置为1是为了第一次dp[i-2]的值
	dp := make([]int, len(s)+1)
	dp[0] = 1
	// s从0开始，dp从1开始
	for i := 1; i < len(s)+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func main() {
	s := "0"
	numDecodings(s)
}
