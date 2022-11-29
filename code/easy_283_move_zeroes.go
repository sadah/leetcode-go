/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func moveZeroes(nums []int) {
	cursor := 0
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[cursor] = nums[cursor], nums[i]
			cursor++
		}
	}
}

// @lc code=end
