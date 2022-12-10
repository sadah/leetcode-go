/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func isPerfectSquare(num int) bool {
	// https://ja.wikipedia.org/wiki/%E3%83%8B%E3%83%A5%E3%83%BC%E3%83%88%E3%83%B3%E6%B3%95
	if num < 2 {
		return true
	}

	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}

	return x*x == num
}

// @lc code=end
