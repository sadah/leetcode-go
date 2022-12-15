/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
package code

import (
	"fmt"
	"math/bits"
)

// @lc code=start
//
// https://leetcode.com/problems/binary-watch/solutions/1617174/go-solution-using-brute-forcing-and-bit-manipulations-beats-100/
//
//lint:ignore U1000 //
func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	if turnedOn > 8 {
		return result
	}

	minutes, hours := uint64((1<<6)-1), uint64(((1<<4)-1)<<6)
	var i uint64
	for i = 0; i < (1<<10)-1; i++ {
		if bits.OnesCount64(i) == turnedOn {
			m, h := i&minutes, i&hours>>6
			if m < 60 && h < 12 {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}

// @lc code=end
