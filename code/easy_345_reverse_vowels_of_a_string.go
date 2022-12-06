/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

package code

import (
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func reverseVowels(s string) string {
	chars := strings.Split(s, "")
	rcs := make([]string, len(chars))
	copy(rcs, chars)
	for i := 0; i < len(rcs)/2; i++ {
		rcs[i], rcs[len(rcs)-i-1] = rcs[len(rcs)-i-1], rcs[i]
	}

	rv := []string{}
	for _, v := range rcs {
		switch v {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			rv = append(rv, v)
		}
	}

	j := 0
	for i, v := range chars {
		switch v {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			chars[i] = rv[j]
			j++
		}
	}
	return strings.Join(chars, "")
}

// @lc code=end
