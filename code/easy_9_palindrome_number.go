/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package code

import "strconv"

// @lc code=start
//
//lint:ignore U1000 //
func isPalindrome1(x int) bool {
	s := strconv.Itoa(x)
	rs := reverse1(s)
	if s == rs {
		return true
	} else {
		return false
	}
}

func reverse1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end
