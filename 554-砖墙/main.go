package main

/*
1. 墙的高度是定长
2. 计算每一行除最右砖块外到最左边的距离
3. 对于每一行而言每块砖到最左边的距离是不同的
4. 所以求的到最左边距离相等的行越多越好
*/

func leastBricks(wall [][]int) int {
	leftWallLength := map[int]int{}
	for _, v :=range wall{
		leftLength := 0
		for _, brick:= range v[:len(v)-1]{
			leftLength += brick
			leftWallLength[leftLength]++
		}
	}
	maxSame := 0
	for _, v := range leftWallLength{
		if maxSame < v{
			maxSame = v
		}
	}
	return len(wall) - maxSame
}
