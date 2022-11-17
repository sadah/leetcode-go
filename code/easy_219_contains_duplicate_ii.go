/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func containsNearbyDuplicate(nums []int, k int) bool {
	// https://leetcode.com/problems/contains-duplicate-ii/discuss/828158
	// key: number
	// value: last position on the left hand side
	lastPosition := make(map[int]int)

	for idx, curNum := range nums {

		lastPos, exist := lastPosition[curNum]

		// Do we have duplicate numbers within distance k ?
		if exist && (idx-lastPos) <= k {
			return true
		}

		// update last position of current number
		lastPosition[curNum] = idx
	}

	return false

}

// @lc code=end
