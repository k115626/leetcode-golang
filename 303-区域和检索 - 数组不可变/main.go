package main

// 这个方法要调用10 ** 4次，如果每次都要计算的话，会超时
// 前缀和
// 某个区域的值未end的前缀和 - start的前缀和
// 统计一次，可以把数值存下来，以后的每次计算都是 O(1)
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {

}
