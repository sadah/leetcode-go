/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func firstUniqChar(s string) int {
	for i, c := range s {
		if strings.Count(s, string(c)) == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
