/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func longestPalindrome(s string) int {
	ret := 0
	ss := strings.Split(s, "")
	sm := map[string]int{}
	for _, str := range ss {
		sm[str] += 1
	}
	hasOdd := false
	for _, v := range sm {
		if v%2 == 1 {
			hasOdd = true
			ret += v - 1
		}else{
			ret += v
		}
	}
	if hasOdd {
		ret += 1
	}
	return ret
}

// @lc code=end
