/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end
