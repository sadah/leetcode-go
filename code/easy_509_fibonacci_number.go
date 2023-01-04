/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func fib(n int) int {
	n0, n1 := 0, 1
	if n == 0 {
		return n0
	} else if n == 1 {
		return n1
	} else if n == 2 {
		return n0 + n1
	}
	for i := 3; i <= n; i++ {
		tmp := n0 + n1
		n0 = n1
		n1 = tmp
	}
	return n0 + n1
}

// @lc code=end
