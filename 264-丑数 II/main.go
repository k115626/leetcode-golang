package main

import "fmt"

/*
dp = [ x, y, z ]
z 的下一个丑数是x,y,z中的还未全部与2，3，5相乘的那个数分别和2，3，5中未乘的数相乘，取最小值
动态规划： 可以保留每个已经计算出来且有序的丑数
三个指针： 可以知道2乘到了哪个丑数，3乘到了哪个丑数，5乘到了哪个丑数
每个丑数只能与这三个数做一次乘法
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i < n+1; i++ {
		num2, num3, num5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(num2, num3), num5)
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
		if dp[i] == num5 {
			p5++
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	n := 30
	res := nthUglyNumber(n)
	fmt.Println(res)
}
