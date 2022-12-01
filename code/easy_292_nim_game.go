/*
 * @lc app=leetcode id=292 lang=golang
 *
 * [292] Nim Game
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end
