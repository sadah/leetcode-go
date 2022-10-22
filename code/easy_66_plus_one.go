/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package code

import (
	"math/big"
	"strconv"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func plusOne(digits []int) []int {
	s := ""
	for _, i := range digits {
		s += strconv.Itoa(i)
	}
	num, _ := new(big.Int).SetString(s, 10)
	num.Add(num, big.NewInt(1))
	s = num.String()
	ss := strings.Split(s, "")
	var ret []int
	for _, sd := range ss {
		id, _ := strconv.Atoi(sd)
		ret = append(ret, id)
	}

	return ret
}

// @lc code=end
