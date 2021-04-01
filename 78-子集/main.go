package main

import "fmt"

// 数组中的数只有两种状态，在子集和不在子集中
// 用1表示在子集中，0表示不在子集中
// 相当于位数等于数组长度的二进制所表示的所有数
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	// 最大值是111，则等同与<1000
	for i := 0; i < 1<<n; i++ {
		subRes := []int{}
		for k, v := range nums {
			if i>>k&1 > 0 {
				subRes = append(subRes, v)
			}
		}
		res = append(res, subRes)
	}
	return res
}

// 迭代法
// 5个元素的所有子集可以看作是 4个元素的所有子集加上第5个元素 然和合并到4个元素的所有子集合中
func subsets1(nums []int) [][]int {
	res := [][]int{[]int{}} // 只包含空集合的集合
	for _, v := range nums {
		addArr := [][]int{}
		for _, n := range res {
			n = append(n, v)
			addArr = append(addArr, n)
		}
		res = append(res, addArr...)
	}

	return res
}

// 回溯法
func subsets2(nums []int) [][]int {
	path, res := []int{}, [][]int{}
	for i := 0; i < len(nums); i++ {
		dfs(nums, i, 0, path, &res)
	}
	return res
}

func dfs(nums []int, k, start int, path []int, res *[][]int) {
	if len(path) == k {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*res = append(*res, pathCopy)
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, k, i+1, path, res)
		path = path[:len(path)-1]
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := subsets2(nums)
	fmt.Println(res)
}
