/*
 * @lc app=leetcode id=561 lang=golang
 *
 * [561] Array Partition
 */

package code

import (
	"sort"
)

// @lc code=start
//
//lint:ignore U1000 //
func arrayPairSum(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		ret += nums[i*2]
	}
	return ret
}

// @lc code=end
