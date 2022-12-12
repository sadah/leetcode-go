/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func canConstruct(ransomNote string, magazine string) bool {
	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}
	}
	return true
}

// @lc code=end
