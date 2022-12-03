/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

package code

// @lc code=start
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return true
}

// @lc code=end
