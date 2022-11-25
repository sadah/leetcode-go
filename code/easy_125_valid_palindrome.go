/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package code

import (
	"regexp"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func isPalindrome2(s string) bool {
	ls := strings.ToLower(s)
	r := regexp.MustCompile("[^0-9a-z]")
	rls := r.ReplaceAllString(ls, "")
	return rls == reverse(rls)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end
