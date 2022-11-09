/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, n := range nums {
		if !m[n] {
			m[n] = true
		} else {
			delete(m, n)
		}
	}
	ret := 0
	for k := range m {
		ret = k
	}
	return ret
}

// @lc code=end
