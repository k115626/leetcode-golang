package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	num1P, num2P := m-1, n-1
	for endP := m + n - 1; num1P >= 0 || num2P >= 0; endP-- {
		if num1P == -1 {
			nums1[endP] = nums2[num2P]
			num2P--
		} else if num2P == -1 {
			nums1[endP] = nums1[num1P]
			num1P--
		} else if nums1[num1P] > nums2[num2P] {
			nums1[endP] = nums1[num1P]
			num1P--
		} else {
			nums1[endP] = nums2[num2P]
			num2P--
		}
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
