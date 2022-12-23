/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

package code

import (
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func repeatedSubstringPattern(s string) bool {
	for i := 0; i < len(s); i++ {
		str := s[0:i]
		ss := strings.Split(s, str)
		if len(strings.Join(ss, "")) == 0 {
			return true
		}
	}
	return false
}

// @lc code=end
