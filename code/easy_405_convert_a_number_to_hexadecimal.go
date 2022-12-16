/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

package code

import (
	"math"
	"strconv"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = int(math.Pow(16,8)) + num
	}
	ints := []int{}
	for num != 0 {
		ints = append(ints, num%16)
		num = num / 16
	}
	ret := []string{}
	for i := len(ints) - 1; i >= 0; i-- {
		switch ints[i] {
		case 10:
			ret = append(ret, "a")
		case 11:
			ret = append(ret, "b")
		case 12:
			ret = append(ret, "c")
		case 13:
			ret = append(ret, "d")
		case 14:
			ret = append(ret, "e")
		case 15:
			ret = append(ret, "f")
		default:
			ret = append(ret, strconv.Itoa(ints[i]))
		}
	}
	return strings.Join(ret, "")
}

// @lc code=end
