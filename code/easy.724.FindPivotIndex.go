/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

package code

// @lc code=start
func pivotIndex(nums []int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		l := nums[0:i]
		lSum := 0
		for _, le := range l {
			lSum += le
		}

		r := nums[i+1 : len]
		rSum := 0
		for _, re := range r {
			rSum += re
		}
		if lSum == rSum {
			return i
		}
	}
	return -1
}

// @lc code=end
