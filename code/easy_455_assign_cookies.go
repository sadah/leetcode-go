/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

package code

import "sort"

// @lc code=start
//
//lint:ignore U1000 //
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	for _, v := range s {
		if ret >= len(g) {
			break
		}

		if v >= g[ret] {
			ret++
		}
	}

	return ret
}

// @lc code=end
