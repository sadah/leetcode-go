/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func wordPattern(pattern string, s string) bool {
	patterns := strings.Split(pattern, "")
	strs := strings.Split(s, " ")
	if len(patterns) != len(strs) {
		return false
	}

	m1 := map[string]string{}
	m2 := map[string]string{}
	for i := 0; i < len(patterns); i++ {
		v, ok := m1[patterns[i]]
		if ok && v == strs[i] {
			continue
		} else if ok && v != strs[i] {
			return false
		} else if !ok {
			m1[patterns[i]] = strs[i]
			m2[strs[i]] = patterns[i]
		}
	}
	return len(m1) == len(m2)
}

// @lc code=end
