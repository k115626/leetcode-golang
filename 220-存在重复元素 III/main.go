package main

import "fmt"

/*
桶排序
1. 桶可以存一组数据
2. 判断两数差值是不是在t内（包含t），将两数分别除以t+1，商相同的差值就在t内
3. 将商相同的存入一个桶，如果一个桶的数据有两个，就返回True
4. 比较左右两个桶的数据，差值在t内的返回True
5. 将超出界限的那个桶子移除，生成新的桶子
6. 边界值：
	0-10 设为0号桶
	-11 - -1 设为-1号桶
*/

// 获取桶子的编号
func getIndex(n, size int) int {
	if n >= 0 {
		return n / size
	}
	return n/size - 1
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	size := t + 1
	mapInfo := map[int]int{}
	for i, v := range nums {
		index := getIndex(v, size)
		if _, ok := mapInfo[index]; ok {
			return true
		}
		if x, ok := mapInfo[index-1]; ok && abs(x-v) <= t {
			return true
		}
		if x, ok := mapInfo[index+1]; ok && abs(x-v) <= t {
			return true
		}
		mapInfo[index] = v
		if i >= k {
			delete(mapInfo, getIndex(nums[i-k], size))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	k, t := 2, 3
	res := containsNearbyAlmostDuplicate(nums, k, t)
	fmt.Println(res)
}
