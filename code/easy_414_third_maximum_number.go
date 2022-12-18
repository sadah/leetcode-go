/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

package code

import (
	"sort"
)

// @lc code=start
//
//lint:ignore U1000 //
func thirdMax(nums []int) int {
	sort.Ints(nums)
	uNums := uniqueInts(nums)
	if len(uNums) < 3 {
		return uNums[len(uNums)-1]
	} else {
		return uNums[len(uNums)-3]
	}
}

func uniqueInts(s []int) []int {
	m := make(map[int]bool)
	var ret []int
	for _, str := range s {
		if _, ok := m[str]; !ok {
			m[str] = true
			ret = append(ret, str)
		}
	}
	return ret
}

// @lc code=end
