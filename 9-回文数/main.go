package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	switch {
	case x < 0:
		return false
	case x < 10:
		return true
	}
	xTostring := strconv.Itoa(x)
	for i := 0; i < len(xTostring)/2; i++ {
		if xTostring[i] != xTostring[len(xTostring)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	x := 121
	res := isPalindrome(x)
	fmt.Println(res)
}
