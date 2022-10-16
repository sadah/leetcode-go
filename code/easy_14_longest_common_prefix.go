/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func longestCommonPrefix(strs []string) string {
	cs := strings.Split(strs[0], "")
	ret := ""
	for _, c := range cs {
		for _, s := range strs {
			if !strings.HasPrefix(s, ret + c) {
				return ret
			}
		}
		ret = ret + c
	}
	return ret
}

// @lc code=end
