/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n = n / 2
	}

	return true
}

// @lc code=end
