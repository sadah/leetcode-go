/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
package code

import (
	"math"
)

// @lc code=start
//
//lint:ignore U1000 //
func maxProfit(prices []int) int {
	min := math.MaxInt64
	mp := 0
	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > mp {
			mp = p - min
		}
		// fmt.Println(mp, min, p)
	}
	return mp
}

// @lc code=end

//lint:ignore U1000 //
func bruteForce(prices []int) int {
	mp := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			p := prices[j] - prices[i]
			if p > mp {
				mp = p
			}
		}
	}
	return mp
}
