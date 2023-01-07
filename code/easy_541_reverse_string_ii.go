/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func reverseStr(s string, k int) string {
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	n := len(s) / k
	if len(s)%k != 0 {
		n++
	}
	reverse("")
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		p := i * k
		q := (i + 1) * k
		if q >= len(s) {
			q = len(s)
		}
		if i%2 != 0 {
			ret = append(ret, s[p:q])
		} else {
			ret = append(ret, reverse(s[p:q]))
		}
	}
	return strings.Join(ret, "")
}

// @lc code=end
