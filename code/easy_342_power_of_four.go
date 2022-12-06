/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 != 0 {
		return false
	}
	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}
	return true
}

// @lc code=end
