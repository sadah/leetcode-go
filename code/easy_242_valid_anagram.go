/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package code

import (
	"sort"
	"strings"
)

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(ts)
	for i := 0; i < len(s); i++ {
		if ss[i] != ts[i] {
			return false
		}
	}
	return true
}

// @lc code=end
