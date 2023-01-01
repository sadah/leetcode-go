/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

package code

import "strconv"

// @lc code=start
//
//lint:ignore U1000 //
func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

// @lc code=end
