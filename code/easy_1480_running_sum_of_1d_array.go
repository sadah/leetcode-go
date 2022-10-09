/*
 * @lc app=leetcode id=1480 lang=golang
 *
 * [1480] Running Sum of 1d Array
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func runningSum(nums []int) []int {
	ret := []int{}
	sum := 0
	for _, num := range nums {
		sum += num
		ret = append(ret, sum)
	}
	return ret
}

// @lc code=end
