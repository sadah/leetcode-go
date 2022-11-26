/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func addDigits(num int) int {
	ret := 0
	for num > 0 {
		ret += num % 10
		num = num / 10

		if num == 0 && ret >= 10 {
			num = ret
			ret = 0
		}
	}
	return ret
}

// @lc code=end
