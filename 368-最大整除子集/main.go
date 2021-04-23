package main

import (
	"fmt"
	"sort"
)

// DP
// 同一个数的不同倍数之间并不存在倍数关系
// 9 18 27     18和27之间不存在倍数关系
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	// f[i] 为i之前的数字，并且以i为结尾的 [整除数组] 的最长长度
	f := make([]int, len(nums))
	// g[i] 为记录f[i]是由哪个下标的状态转移过来的, 即 f[i] = f[g[i]] +1
	g := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		maxLength := 1
		index := i
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if f[j]+1 > maxLength {
					maxLength = f[j] + 1
					index = j
				}
			}
		}
		f[i] = maxLength
		g[i] = index
	}
	max := 0
	idMax := 0
	for i := 0; i < len(f); i++ {
		if f[i] > max {
			max = f[i]
			idMax = i
		}
	}
	ans := []int{}
	/*
		1. f 中存的是每个位置对应的最大整数
		2. g 中存的是每个位置对应的上个位置的索引
		nums: 【9, 18, 54, 90, 108, 180, 360, 540, 720】
				11	9		7		 5	 3			1
		f:	  [1   2    3   3   4    4    5    5    6]
					10		8		6	 4			2
		g:   [0   0    1   1   2    3    5    4   6]
	*/
	for len(ans) < max {
		ans = append(ans, nums[idMax])
		idMax = g[idMax]
	}
	return ans
}

func main() {
	nums := []int{720, 9, 18, 54, 90, 108, 180, 360, 540}
	res := largestDivisibleSubset(nums)
	fmt.Println(res)
}
