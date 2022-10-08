/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// @lc code=end
