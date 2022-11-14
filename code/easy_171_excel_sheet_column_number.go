/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

package code

import (
	"math"
)

// @lc code=start
//
//lint:ignore U1000 //
func titleToNumber(columnTitle string) int {
	chars := []byte(columnTitle)
	// reverse order
	for h, t := 0, len(chars)-1; h < t; h, t = h+1, t-1 {
		chars[h], chars[t] = chars[t], chars[h]
	}
	ret := 0
	for i, c := range chars {
		ret += (int(c) - 64) * int(math.Pow(26, float64(i)))
	}
	return ret
}

// @lc code=end
