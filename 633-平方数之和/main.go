package main

import "math"

/*
1. 两数取值范围[0, 根号c]
2. 遍历枚举
3. 双指针
*/
func judgeSquareSum(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))
	for a <= b {
		if a*a+b*b == c {
			return true
		} else if a*a+b*b > c {
			b--
		} else {
			a++
		}
	}
	return false
}
