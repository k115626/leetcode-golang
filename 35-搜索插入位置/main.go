package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n
	ans := n
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			ans = mid
			right = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {

}
