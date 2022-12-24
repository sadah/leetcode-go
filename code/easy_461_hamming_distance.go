/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

package code

import (
	"fmt"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func hammingDistance(x int, y int) int {
	return strings.Count(fmt.Sprintf("%b", x^y), "1")
}

// @lc code=end
