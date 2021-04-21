package main

import "fmt"

// KMP 算法
func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	nextVal := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = nextVal[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		nextVal[i] = j
	}
	for i, j := 0, 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = nextVal[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func main() {
	haystack := "ababaabbbbababbaabaaabaabbaaaabbabaabbbbbbabbaabbabbbabbbbbaaabaababbbaabbbabbbaabbbbaaabbababbabbbabaaabbaabbabababbbaaaaaaababbabaababaabbbbaaabbbabb"
	needle := "abbabbbabaa"
	res := strStr(haystack, needle)
	fmt.Println(res)
}
