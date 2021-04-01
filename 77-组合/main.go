package main

import "fmt"

// 回溯算法 -- 递归中嵌套for循环 递归中嵌套for循环 递归中嵌套for循环
// 递归函数的参数
// 确定中止条件
// 单层递归的逻辑

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || n < k {
		return [][]int{}
	}
	path := []int{}
	res := [][]int{}
	generateCombine(n, k, 1, path, &res)
	return res
}

// n: 数组的size
// k: 结果长度
// start: 选择集从哪个位置开始
// path: 节点值，也是满足条件的一次结果
// res: 返回的二维数组

// 数组是从1到n， 所以start既是值，还是索引+1
func generateCombine(n, k, start int, path []int, res *[][]int) {
	// 中止条件
	if len(path) == k {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*res = append(*res, pathCopy)
		return
	}
	// for循环
	// 里面是单次f递归的逻辑
	// 可选集就是从开始到最后的所有值
	for i := start; i < n+1; i++ {
		path = append(path, i)
		generateCombine(n, k, i+1, path, res)
		path = path[:len(path)-1] // 回溯函数
	}
	return

}

func main() {
	n, k := 5, 3
	res := combine(n, k)
	fmt.Println(res)
}
