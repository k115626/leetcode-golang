package main

/*
1. 船的最小运送能力要大于最大的一件商品的重量
2. 船的最大的运送能力要小于所有的商品运送之和
3. 在次区间使用二分法来求解
*/

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, v := range weights {
		if v > left {
			left = v
		}
		right += v
	}
	for left < right {
		d := 1                        // 运送天数
		mid := (right-left)>>1 + left // 船的装载能力
		curr := 0                     // 当前船的装载重量
		for _, v := range weights {
			if (curr + v) > mid {
				d++
				curr = 0
			}
			curr += v
		}
		if d > D {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {

}
