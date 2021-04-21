package main

import (
	"fmt"
)

// 暴力求解
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			// need串的长度已经大于原串剩余比较的长度
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

func main() {
	haystack := "hello"
	needle := "ll"
	res := strStr(haystack, needle)
	fmt.Println(res)
}
