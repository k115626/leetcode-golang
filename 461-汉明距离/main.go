package main

import "fmt"

func hammingDistance(x int, y int) int {
	// 1 0001
	// 4 0100
	// 5 0101
	//s := x ^ y
	// 解法一
	//return bits.OnesCount(uint(s))
	// 解法二
	ans := 0
	//for s > 0 {
	//	ans += s & 1
	//	s = s >> 1
	//}
	// 解法三
	// Brian Kernighan 算法
	// 对于任何一个n， 将n与n-1做&运算，得到的数是将n的最右侧的1位置为0的结果
	// 对于题目，是求两个数字异或后得到数值的位置为1的个数
	//for s > 0 {
	//	ans += 1
	//	s = s & (s - 1)
	//}
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return ans
}

func main() {
	x, y := 1, 4
	ans := hammingDistance(x, y)
	fmt.Println(ans)
}
