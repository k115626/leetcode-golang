package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	factors := []int{2, 3, 5}
	for _, v := range factors {
		for n%v == 0 {
			n /= v
		}
	}
	return n == 1
}

func main() {
	n := 2
	res := isUgly(n)
	fmt.Println(res)
}
