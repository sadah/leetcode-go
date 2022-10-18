/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func removeDuplicates(nums []int) int {
	index := 1

	if len(nums) == 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[index] = nums[i]
			index += 1
		}
	}
	return index
}

// @lc code=end
