package main

/*
1. 每个子组的最小值是数组中的最大值
2. 每个子组的最大值是数组和
*/
func splitArray(nums []int, m int) int {
	min, max := 0, 0
	for _, v := range nums {
		if v > min {
			min = v
		}
		max += v
	}
	// 二分查找
	for min < max {
		mid := (max-min)>>1 + min // 数组的容量
		group := 1                // 分组数
		curr := 0                 // 当前组的数量
		for _, v := range nums {
			if curr+v > mid {
				group++
				curr = 0
			}
			curr += v
		}
		if group > m {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func main() {

}
