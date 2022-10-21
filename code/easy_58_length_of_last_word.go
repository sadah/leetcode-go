/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func lengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")
	ret := ""
	for _, w := range ss {
		if w != "" {
			ret = w
		}
	}
	return len(ret)
}

// @lc code=end
