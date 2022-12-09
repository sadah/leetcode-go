/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func intersect(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	for _, n := range nums1 {
		m1[n]++
	}
	m2 := map[int]int{}
	for _, n := range nums2 {
		m2[n]++
	}
	ret := []int{}
	for k, v := range m1 {
		j := 0
		if v <= m2[k] {
			j = v
		} else {
			j = m2[k]
		}
		for i := 0; i < j; i++ {
			ret = append(ret, k)
		}
	}
	return ret
}

// @lc code=end
