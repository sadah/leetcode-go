/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func countSegments(s string) int {
	chars := strings.Split(s, " ")
	ret := 0
	for _, c := range chars {
		if c != "" {
			ret++
		}
	}
	return ret
}

// @lc code=end
