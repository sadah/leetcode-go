/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */
package code

import (
	"strconv"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func addStrings(num1 string, num2 string) string {
	ret := []string{}
	rs1 := strings.Split(num1, "")
	rs2 := strings.Split(num2, "")
	l1 := len(rs1) - 1
	l2 := len(rs2) - 1
	tmp := 0
	for l1 >= 0 || l2 >= 0 {
		n1, n2 := 0, 0
		if l1 >= 0 {
			n1, _ = strconv.Atoi(rs1[l1])
		}
		if l2 >= 0 {
			n2, _ = strconv.Atoi(rs2[l2])
		}
		ret = append(ret, strconv.Itoa((n1+n2+tmp)%10))
		tmp = (n1 + n2 + tmp) / 10
		l1--
		l2--
	}
	if tmp != 0 {
		ret = append(ret, strconv.Itoa(tmp))
	}
	return reverse2(strings.Join(ret, ""))
}

func reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end
