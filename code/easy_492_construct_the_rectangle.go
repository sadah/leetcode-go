/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */
package code

import "math"

// @lc code=start
//
//lint:ignore U1000 //
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w != 0 {
		w--
	}
	return []int{area / w, w}
}

// @lc code=end
