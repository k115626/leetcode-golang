package main

import "fmt"

var mapInfo = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		if mapInfo[s[i+1]] > mapInfo[s[i]] {
			sum -= mapInfo[s[i]]
		} else {
			sum += mapInfo[s[i]]
		}
	}
	return sum + mapInfo[s[len(s)-1]]
}

func main() {
	s := "MCMXCIV"
	res := romanToInt(s)
	fmt.Println(res)
}
