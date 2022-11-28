/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

package code

import "sort"

// @lc code=start
//
//lint:ignore U1000 //
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, n := range nums {
		if n != i {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
