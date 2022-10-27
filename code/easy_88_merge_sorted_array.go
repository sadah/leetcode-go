/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
package code

import (
	"sort"
)

// @lc code=start
//
//lint:ignore U1000 //
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	nums1 = append(nums1, nums2...)
	sort.Sort(sort.IntSlice(nums1))
}

// @lc code=end

// func merge2(nums1 []int, m int, nums2 []int, n int) {
// 	cnums1 := make([]int, m+n)
// 	copy(cnums1, nums1[:])

// 	cnums1 = cnums1[0:m]
// 	if m == 0 {
// 		nums1[0] = nums2[0]
// 	}
// 	if n == 0 {
// 		return
// 	}
// 	if m == 1 && n == 1 {
// 		if nums1[0] < nums2[0] {
// 			nums1[1] = nums2[0]
// 		} else {
// 			tmp := nums1[1]
// 			nums1[0] = nums2[0]
// 			nums1[1] = tmp
// 		}
// 	}
// 	i, j, k := 0, 0, 0
// 	for m-1 != 0 || n-1 != 0 {
// 		if m != 0 && cnums1[i] <= nums2[j] {
// 			nums1[k] = cnums1[i]
// 			i++
// 			k++
// 			m--
// 		} else if n != 0 {
// 			nums1[k] = nums2[j]
// 			j++
// 			k++
// 			n--
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Println(nums1)
// }
