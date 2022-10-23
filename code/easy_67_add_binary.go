/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

package code

import (
	"math/big"
	"strconv"
)

// @lc code=start
//
//lint:ignore U1000 //
func addBinary(a string, b string) string {
	i1, _ := new(big.Int).SetString(a, 2)
	i2, _ := new(big.Int).SetString(b, 2)
	i3 := new(big.Int).Add(i1, i2)
	return i3.Text(2)
}

// @lc code=end
