/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// @lc code=end
