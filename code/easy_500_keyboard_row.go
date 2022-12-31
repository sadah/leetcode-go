/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func findWords(words []string) []string {
	chars := make(map[rune]int, 26)
	for _, c := range "qwertyuiop" {
		chars[c] = 1
	}
	for _, c := range "asdfghjkl" {
		chars[c] = 2
	}
	for _, c := range "zxcvbnm" {
		chars[c] = 3
	}
	ret := []string{}

	for _, word := range words {
		lword := strings.ToLower(word)
		same := true
		for _, c := range lword {
			if chars[rune(lword[0])] != chars[c] {
				same = false
				break
			}
		}
		if same {
			ret = append(ret, word)
		}
	}
	return ret
}

// @lc code=end
