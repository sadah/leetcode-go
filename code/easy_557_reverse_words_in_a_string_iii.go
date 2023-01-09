/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func reverseWords(s string) string {
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	words := strings.Split(s, " ")
	ret := make([]string, len(words))
	for i, word := range words {
		ret[i] = reverse(word)
	}
	return strings.Join(ret, " ")
}

// @lc code=end
