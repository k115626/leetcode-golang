package main

import "fmt"

/*
解析情况:
1. 只有一间房： 选a
2. 两家房： 选a,b最大值
3. 大于两间:
	分为连个序列，选第一间 && 不选最后一间
                不选第一件 && 选最后一间
*/
func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[1:len(nums)]), _rob(nums[:len(nums)-1]))
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	nums := []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
