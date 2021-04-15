package main

import "fmt"

/*
动态规划：
a, b, c 三间房
1. 有一间房，选a
2. 有两间房，选a,b的最大值
3. 大于两间房
	1） 选c房，则为c + c-2 的和
	2） 不选c，则为c-1 的值
dp[i] 表示i间房可偷窃的最大值
动态方程： dp[i] = max(dp[i-2]+nums[i], dp[i-1])
边界： dp[i] = {

*/
func rob(nums []int) int {
	dp := make([]int, len(nums))
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		dp[1] = max(nums[0], nums[1])
	}
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
