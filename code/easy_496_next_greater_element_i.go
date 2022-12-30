/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	j := 0

	for i := 0; i < len(nums1); i++ {
		found := false
		for j = 0; j < len(nums2); j++ {
			if found && nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}

			if nums2[j] == nums1[i] {
				found = true
			}
		}
		if j == len(nums2) {
			res[i] = -1
		}
	}

	return res
}

// @lc code=end
