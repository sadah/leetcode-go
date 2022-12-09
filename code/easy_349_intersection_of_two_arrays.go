/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
package code

// @lc code=start
//
//lint:ignore U1000 //
func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}
	for _, n := range nums1 {
		m1[n] = true
	}
	m2 := map[int]bool{}
	for _, n := range nums2 {
		m2[n] = true
	}
	ret := []int{}
	for k := range m1 {
		_, ok := m2[k]
		if ok {
			ret = append(ret, k)
		}
	}
	return ret
}

// @lc code=end
