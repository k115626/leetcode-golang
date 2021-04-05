package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	container := make([][]byte, numRows)
	currRow, addRow := 0, 1
	for _, v := range s {
		if currRow == numRows-1 {
			addRow = -1
		} else if currRow == 0 {
			addRow = 1
		}
		container[currRow] = append(container[currRow], byte(v))
		currRow += addRow
	}
	res := make([]byte, 0)
	for _, v := range container {
		res = append(res, v...)
	}
	return string(res)
}

func main() {
	s := "PAC"
	numRows := 2
	res := convert(s, numRows)
	fmt.Println(res)
}
