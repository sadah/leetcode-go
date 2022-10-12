/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func search(nums []int, target int) int {
	l := len(nums)
	start := 0
	end := l - 1
	mid := 0
	for {
		mid = (start + end) / 2
		if end < start {
			return -1
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
}

// @lc code=end
