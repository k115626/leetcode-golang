package main

import "fmt"

// 栈的数据结构特性
func isValid(s string) bool {
	stack := []rune{'?'}
	mapInfo := map[rune]rune{
		'{': '}', '[': ']', '(': ')', '?': '?',
	}
	for _, v := range s {
		if _, ok := mapInfo[v]; ok {
			stack = append(stack, v)
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mapInfo[pop] != v {
				return false
			}
		}
	}
	return len(stack) == 1
}

func main() {
	s := "()[]{}"
	res := isValid(s)
	fmt.Println(res)
}
