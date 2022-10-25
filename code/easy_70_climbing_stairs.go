/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package code

// @lc code=start
//
//lint:ignore //U1000
func climbStairs(n int) int {
	if n == 1 {
		return 1;
	}
	first := 1;
	second := 2;
	for i := 3; i <= n; i++ {
		third := first + second;
		first = second;
		second = third;
	}
	return second;
}


// @lc code=end
