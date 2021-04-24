package main

/*
截止到当前位置的最大连续数组之和
 = 当前位置的所有数据之和 - 该位置之前最小的数组之和
*/
func maxSubArray(nums []int) int {
	minNums := 0       // 记录最小和
	currNums := 0      // 当前总和
	maxNums := -100000 // 记录最大结果
	for _, v := range nums {
		currNums += v
		maxNums = max(maxNums, currNums-minNums)
		minNums = min(minNums, currNums)
	}
	return maxNums
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func min(n, m int) int {
	if n > m {
		return m
	}
	return n
}

func main() {
	nums := []int{5, 4, -1, 7, 8}
	maxSubArray(nums)
}
