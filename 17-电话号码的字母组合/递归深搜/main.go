package main

import "fmt"

// "23" 中的 '2' 是byte类型，数值是50
// 使用 '2' - '0' 可以拿到int

// 首先需要在全局定义 res 变量
// 然后在初始函数时将这个变量的内容初始化，否则res中还保存有之前的执行结果

var (
	mapInfo = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res     = []string{}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = []string{} // 不加这一行，res会将之前存入的数据也保存下来
	dfs(digits, 0, "")
	return res
}

func dfs(digits string, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	num := digits[index]
	letter := mapInfo[num-'0']
	for i := 0; i < len(letter); i++ {
		dfs(digits, index+1, s+string(letter[i]))
	}
}

func main() {
	digits := "34"
	res := letterCombinations(digits)
	fmt.Println(res)
}
