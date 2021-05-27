package main

/*
1. 使用map统计每个数字出现的次数，遍历map获取出现一次的数字，使用map需要额外的空间
2. 进阶
	1. 将每个数字转化为二进制数
	2. 重复的数字只能出现三次
	3. 重复的数字对应的二进制位相加，必被3整除
*/
//func singleNumber(nums []int) int {
//	freq := map[int]int{}
//	for _, v := range nums {
//		freq[v]++
//	}
//	for num, occ := range freq {
//		if occ == 1 {
//			return num
//		}
//	}
//	return 0
//}

func singleNumber(nums []int) int {
	ans := int32(0)
	// nums[i] 属于int32内
	for i := 0; i < 32; i++ {
		// 每一个二进制位的和值
		total := int32(0)
		for _, v := range nums {
			total += int32(v) >> i & 1
		}
		if total%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
