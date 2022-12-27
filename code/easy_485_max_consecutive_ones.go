/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func findMaxConsecutiveOnes(nums []int) int {
	max, tmp := 0, 0
	for _, n := range nums {
		if n == 1 {
			tmp += 1
			if tmp > max {
				max = tmp
			}
		} else {
			tmp = 0
		}
	}
	return max
}

// @lc code=end
