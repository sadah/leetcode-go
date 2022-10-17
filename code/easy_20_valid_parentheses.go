/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package code

import "strings"

// @lc code=start
//
//lint:ignore U1000 //
func isValid(s string) bool {
	cs := strings.Split(s, "")
	ss := []string{}
	for _, c := range cs {
		switch c {
		case "(":
			ss = append(ss, ")")
		case "[":
			ss = append(ss, "]")
		case "{":
			ss = append(ss, "}")
		case ")":
			if len(ss) == 0 {
				return false
			} else if ss[len(ss)-1] != ")" {
				return false
			}
			ss = ss[0 : len(ss)-1]
		case "]":
			if len(ss) == 0 {
				return false
			} else if ss[len(ss)-1] != "]" {
				return false
			}
			ss = ss[0 : len(ss)-1]
		case "}":
			if len(ss) == 0 {
				return false
			} else if ss[len(ss)-1] != "}" {
				return false
			}
			ss = ss[0 : len(ss)-1]
		}
	}
	if len(ss) != 0 {
		return false
	} else {
		return true
	}
}

// @lc code=end
