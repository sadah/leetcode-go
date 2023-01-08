/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func checkRecord(s string) bool {
	total_a := 0
	late_cnt := 0
	for _, c := range s {
		switch string(c) {
		case "A":
			total_a++
			late_cnt = 0
		case "P":
			late_cnt = 0
		case "L":
			late_cnt++
		}
		if total_a >= 2 || late_cnt >= 3 {
			return false
		}
	}
	return true
}

// @lc code=end
