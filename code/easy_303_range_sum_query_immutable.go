/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

package code

// @lc code=start
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

//lint:ignore ST1006 //
func (this *NumArray) SumRange(left int, right int) int {
	tmp := this.nums[left : right+1]
	ret := 0
	for _, n := range tmp {
		ret += n
	}
	return ret
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
