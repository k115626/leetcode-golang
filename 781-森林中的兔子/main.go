package main

import "fmt"

func numRabbits(answers []int) int {
	mapInfo := make(map[int]int)
	for _, v := range answers {
		mapInfo[v]++
	}
	ans := 0
	for k, v := range mapInfo {
		ans += (v + k) / (k + 1) * (k + 1)
	}
	return ans
}

func main() {
	answers := []int{0, 1, 0, 1, 0}
	res := numRabbits(answers)
	fmt.Println(res)
}
