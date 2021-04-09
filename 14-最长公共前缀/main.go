package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		prefix = twoString(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func twoString(s1, s2 string) string {
	n1, n2 := len(s1), len(s2)
	min := 0
	if n1 > n2 {
		min = n2
	} else {
		min = n1
	}
	index := 0
	for index < min && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
