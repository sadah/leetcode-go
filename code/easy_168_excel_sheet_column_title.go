/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func convertToTitle(columnNumber int) string {
	chars := []byte{}

	for columnNumber > 0 {
		columnNumber--
		chars = append(chars, byte('A'+columnNumber%26))
		columnNumber /= 26
	}

	// reverse order
	for h, t := 0, len(chars)-1; h < t; h, t = h+1, t-1 {
		chars[h], chars[t] = chars[t], chars[h]
	}

	return string(chars)
}

// @lc code=end
