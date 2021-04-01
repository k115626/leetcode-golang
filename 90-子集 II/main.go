package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	path, res := []int{}, [][]int{}
	for i := 0; i < len(nums)+1; i++ {
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
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, k, i+1, path, res)
		path = path[:len(path)-1]
	}
	return
}

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
