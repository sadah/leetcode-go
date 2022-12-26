/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

package code

import (
	"strings"
)

// @lc code=start
//
//lint:ignore U1000 //
func licenseKeyFormatting(s string, k int) string {
	str := strings.Replace(strings.ToUpper(s), "-", "", -1)
	cur := len(str) % k
	if cur == 0 {
		cur += k
	}
	for cur < len(str) {
		str = str[:cur] + "-" + str[cur:]
		cur += k + 1
	}
	return str
}

// @lc code=end
