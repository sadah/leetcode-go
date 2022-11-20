/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

package code

import "strconv"

// @lc code=start
//
//lint:ignore U1000 //
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	ret := []string{}
	h := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}
		if h == i {
			ret = append(ret, strconv.Itoa(nums[i]))
		} else {
			tmp := strconv.Itoa(nums[h]) + "->" + strconv.Itoa(nums[i])
			ret = append(ret, tmp)
		}
		h = i + 1
	}
	return ret
}

// @lc code=end
