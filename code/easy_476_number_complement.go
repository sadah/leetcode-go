/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

package code

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
//
//lint:ignore U100 //
func findComplement(num int) int {
	chars := strings.Split(fmt.Sprintf("%b", num), "")
	ret := make([]string, len(chars))
	for i, c := range chars {
		if c == "0" {
			ret[i] = "1"
		} else {
			ret[i] = "0"
		}
	}
	r, _ := strconv.ParseInt(strings.Join(ret, ""), 2, 0)
	return int(r)
}

// @lc code=end
