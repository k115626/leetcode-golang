package main

import (
	"fmt"
)

// 求解运算值的时候可以使用栈的数据结构
func clumsy(N int) int {
	res := 0
	nStack := []int{N}
	index := 0
	N--
	for N > 0 {
		switch index % 4 {
		case 0:
			nStack[len(nStack)-1] *= N
		case 1:
			nStack[len(nStack)-1] /= N
		case 2:
			nStack = append(nStack, N)
		default:
			nStack = append(nStack, -N)
		}
		index++
		N--
	}
	for _, v := range nStack {
		res += v
	}
	return res
}

func main() {
	N := 1
	res := clumsy(N)
	fmt.Println(res)
}
