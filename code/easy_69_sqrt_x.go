/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package code

import "math"

// @lc code=start
//
//lint:ignore U1000 //
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// @lc code=end
