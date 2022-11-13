/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

package code

import "sort"

// @lc code=start
//
//lint:ignore U1000 //
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// @lc code=end
