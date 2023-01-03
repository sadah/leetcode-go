/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func checkPerfectNumber(num int) bool {
	switch num {
	case 6, 28, 496, 8128, 33550336:
		return true
	default:
		return false
	}
}

// @lc code=end
