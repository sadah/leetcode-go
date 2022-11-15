/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

package code

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func reverseBits(num uint32) uint32 {
	s := fmt.Sprintf("%b", num)
	s = fmt.Sprintf("%032s", s)

	chars := strings.Split(s, "")
	for h, t := 0, len(chars)-1; h < t; h, t = h+1, t-1 {
		chars[h], chars[t] = chars[t], chars[h]
	}

	rs := strings.Join(chars, "")
	ret, _ := strconv.ParseUint(rs, 2, 32)

	return uint32(ret)
}

// @lc code=end
