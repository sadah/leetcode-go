/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func getRow(rowIndex int) []int {
	ret := make([][]int, rowIndex+1)
	ret[0] = []int{1}

	for i := 1; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		prev := ret[i-1]
		row[0] = 1
		for j := 1; j < i; j++ {
			row[j] = prev[j-1] + prev[j]
		}
		row[i] = 1
		ret[i] = row
	}
	return ret[rowIndex]
}

// @lc code=end
