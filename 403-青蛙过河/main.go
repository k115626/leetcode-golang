package main

func canCross(stones []int) bool {
	hashMap := map[int]bool{}
	return helper(stones, 0, 0, hashMap)
}

// index: 当前所在的索引
// k: 上一次跳的步数
func helper(stones []int, index, k int, hashMap map[int]bool) bool {
	key := 1000*index + k
	if hashMap[key] {
		return false
	} else {
		hashMap[key] = true
	}
	for i := index + 1; i < len(stones); i++ {
		step := stones[i] - stones[index]
		if step >= k-1 && step <= k+1 {
			if helper(stones, i, step, hashMap) {
				return true
			}
		} else if step > k+1 {
			break
		}
	}
	return index == len(stones)-1
}
