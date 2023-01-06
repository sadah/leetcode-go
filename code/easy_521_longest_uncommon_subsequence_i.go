/*
 * @lc app=leetcode id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	} else if len(a) >= len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

// @lc code=end
