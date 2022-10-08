/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

package code

import (
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func isSubsequence(s string, t string) bool {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	ret := []string{}
	for _, s := range ss {
		for i, t := range ts {
			if t == s {
				ts = ts[i+1:]
				ret = append(ret, t)
				break
			}
		}
	}

	if strings.Join(ret, "") == s {
		return true
	} else {
		return false
	}
}

// @lc code=end
