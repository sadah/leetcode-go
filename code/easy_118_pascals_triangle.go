/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	ret[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		prev := ret[i-1]
		row[0] = 1
		for j := 1; j < i; j++ {
			row[j] = prev[j-1] + prev[j]
		}
		row[i] = 1
		ret[i] = row
	}
	return ret
}

// @lc code=end
