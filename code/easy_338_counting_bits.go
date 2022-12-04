/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

package code

import (
	"fmt"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func countBits(n int) []int {
	ret := []int{}
	for i := 0; i <= n; i++ {
		s := fmt.Sprintf("%b", i)
		cs := strings.Split(s, "")
		r := 0
		for _, c := range cs {
			if c == "1" {
				r += 1
			}
		}
		ret = append(ret, r)
	}
	return ret
}

// @lc code=end
