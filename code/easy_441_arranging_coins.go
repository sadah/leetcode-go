/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func arrangeCoins(n int) int {
	i := 1
	for n >= i {
		n -= i
		i++
	}
	return i - 1
}

// @lc code=end
