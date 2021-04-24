package main

/*
dp[i]: 和为i的组合总数
抛出组合的最后一个数字，则有dp[i] = 所有的dp[i-抛出]组合之和
抛出的数字可能为nums中的任何一个数字，所以dp[i] = len(nums)个dp[i-抛出]组合之和
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

func main() {

}
