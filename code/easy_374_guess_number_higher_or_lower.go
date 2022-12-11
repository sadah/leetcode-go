/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

package code

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
//lint:ignore U1000 //
func guessNumber(n int) int {
	l := 1
	h := n
	for l <= h {
		m := l + (h-l)/2
		ret := guess(m)
		if ret == 0 {
			return m
		} else if ret < 0 {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// @lc code=end

func guess(n int) int {
	return 0
}
