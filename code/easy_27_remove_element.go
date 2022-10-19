/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func removeElement(nums []int, val int) int {
	index := 0
	for _, n := range nums {
		if n != val {
			nums[index] = n
			index++
		}
	}
	return index
}

// @lc code=end
