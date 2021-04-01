package main

import "fmt"

//  定义二维数组存储结果
//  ss = "abcb"
//     a    b    c    b
//    a
//    b
//    c
//    b
// 从 a-a 是不是回文
// 从 a-b 是不是回文
// 从 b-c 是不是回文

func longestPalindrome(s string) string {
	// 定义一个二维数组储存结果
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	var res string
	// lenght 是回文字串的长度
	for lenght := 0; lenght < len(s); lenght++ {
		// start 是字串的起始位置
		for start := 0; start < len(s)-lenght; start++ {
			end := start + lenght
			switch lenght {
			case 0:
				dp[start][end] = true
			case 1:
				dp[start][end] = s[start] == s[end]
			default:
				dp[start][end] = s[start] == s[end] && dp[start+1][end-1]
			}
			if dp[start][end] {
				res = s[start : end+1]
			}
		}
	}
	return res
}

func main() {
	s := "aabcb"
	res := longestPalindrome(s)
	fmt.Println(res)
}
