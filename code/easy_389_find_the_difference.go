/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func findTheDifference(s string, t string) byte {
	for _, c := range t {
		if strings.Count(t, string(c)) != strings.Count(s, string(c)) {
			return byte(c)
		}
	}
	return 0
}

// @lc code=end
