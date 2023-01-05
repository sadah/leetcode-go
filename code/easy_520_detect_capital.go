/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

package code

import (
	"fmt"
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	l := strings.ToLower(word)
	u := strings.ToUpper(word)
	if l == word || u == word {
		fmt.Print(l, u)
		return true
	}
	if word[:1] == strings.ToUpper(word[:1]) && word[1:] == strings.ToLower(word[1:]) {
		return true
	}
	return false
}

// @lc code=end
