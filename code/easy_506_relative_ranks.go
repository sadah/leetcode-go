/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

package code

import (
	"sort"
	"strconv"
)

// @lc code=start
//
//lint:ignore U1000 //
func findRelativeRanks(score []int) []string {
	sorted_score := make([]int, len(score))
	copy(sorted_score, score)
	sort.Ints(sorted_score)
	for i := 0; i < len(sorted_score)/2; i++ {
		sorted_score[i], sorted_score[len(sorted_score)-i-1] = sorted_score[len(sorted_score)-i-1], sorted_score[i]
	}

	m := make(map[int]string, len(score))
	for i, v := range sorted_score {
		if i == 0 {
			m[v] = "Gold Medal"
		} else if i == 1 {
			m[v] = "Silver Medal"
		} else if i == 2 {
			m[v] = "Bronze Medal"
		} else {
			m[v] = strconv.Itoa(i + 1)
		}
	}
	ret := make([]string, len(score))
	for i, v := range score {
		ret[i] = m[v]
	}
	return ret
}

// @lc code=end
