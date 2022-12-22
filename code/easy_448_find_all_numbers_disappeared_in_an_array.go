/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)

	s := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		s[num] = struct{}{}
	}

	ints := make([]int, n)
	for i := 1; i <= n; i++ {
		ints[i-1] = i
	}

	ret := []int{}
	for _, num := range ints {
		if _, ok := s[num]; ok {
			continue
		}
		ret = append(ret, num)
	}

	return ret
}

// @lc code=end
