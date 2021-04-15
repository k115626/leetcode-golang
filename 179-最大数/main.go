package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 采用选择排序的方式，每次获取到最大的一个数值，置换到列表的最左侧
// 比较 a+b 和 b+a 的大小
// a+b 组合更大，则 a 数值左边
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	strNums := []string{}
	for _, v := range nums {
		strNums = append(strNums, strconv.Itoa(v))
	}
	for i := 0; i < len(strNums)-1; i++ {
		for j := i + 1; j < len(strNums); j++ {
			if strNums[i]+strNums[j] < strNums[j]+strNums[i] {
				strNums[i], strNums[j] = strNums[j], strNums[i]
			}
		}
	}
	strRes := strings.Join(strNums, "")
	if v, _ := strconv.Atoi(strRes); v == 0 {
		return "0"
	}
	return strRes
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	res := largestNumber(nums)
	fmt.Println(res)
}
