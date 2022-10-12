/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

package code

func isBadVersion(version int) bool {
	return false
}

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//lint:ignore U1000 //
func firstBadVersion(n int) int {
	start := 1
	end := n
	for {
		mid := (start + end) / 2
		if start == end {
			return start
		}
		if isBadVersion(mid) && isBadVersion(mid+1) {
			end = mid
		}
		if !isBadVersion(mid) && !isBadVersion(mid+1) {
			start = mid
		}
		if !isBadVersion(mid) && isBadVersion(mid+1) {
			return mid + 1
		}
	}
}

// @lc code=end
