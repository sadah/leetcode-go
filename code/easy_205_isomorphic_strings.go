/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

package code

import (
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func isIsomorphic(s string, t string) bool {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	sm := map[string]string{}
	tm := map[string]string{}
	for i := 0; i < len(ss); i++ {
		sv, ok := sm[ss[i]]
		if ok {
			if sv != ts[i] {
				return false
			}
		}

		tv, ok := tm[ts[i]]
		if ok {
			if tv != ss[i] {
				return false
			}
		}
		sm[ss[i]] = ts[i]
		tm[ts[i]] = ss[i]
	}
	return true
}

// @lc code=end
