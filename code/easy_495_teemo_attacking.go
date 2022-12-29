/*
 * @lc app=leetcode id=495 lang=golang
 *
 * [495] Teemo Attacking
 */

package code

import "math"

// @lc code=start
//
//lint:ignore U1000 //
func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		ret += int(math.Min(float64(timeSeries[i+1]-timeSeries[i]), float64(duration)))
	}
	return ret + duration
}

// @lc code=end
