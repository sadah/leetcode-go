/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

package main

import "fmt"

// @lc code=start
func pivotIndex(nums []int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		l := nums[0:i]
		lSum := 0
		for _, le := range l {
			lSum += le
		}

		r := nums[i+1 : len]
		rSum := 0
		for _, re := range r {
			rSum += re
		}
		if lSum == rSum {
			return i
		}
	}
	return -1
}

// @lc code=end

func main() {
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
