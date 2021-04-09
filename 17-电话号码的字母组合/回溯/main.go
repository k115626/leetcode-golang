package main

import "fmt"

var mapInfo = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}
var res = make([]string, 0)

// 回溯算法
// 1. 确定递归的参数
// 2. 递归的终止条件
// 3. 单层递归的逻辑
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = []string{}
	dfs(digits, "")
	return res
}

func dfs(digits, path string) {
	if len(digits) == 0 {
		res = append(res, path)
		return
	}
	// a[0:1]  获取的值是字符串类型
	// a[0]    获取的值是byte类型
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(mapInfo[k]); i++ {
		path += mapInfo[k][i]
		dfs(digits, path)
		path = path[0 : len(path)-1]
	}
}

func main() {
	digits := "345"
	res := letterCombinations(digits)
	fmt.Println(res)
}
