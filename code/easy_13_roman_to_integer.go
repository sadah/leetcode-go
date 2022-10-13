/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	// I             1
	// V             5
	// X             10
	// L             50
	// C             100
	// D             500
	// M             1000
	// IV, IX => v, x => 4, 9
	// XL, XC => l, c => 40, 90
	// CD, CM => d, m => 400, 900
	s = strings.ReplaceAll(s, "IV", "v")
	s = strings.ReplaceAll(s, "IX", "x")
	s = strings.ReplaceAll(s, "XL", "l")
	s = strings.ReplaceAll(s, "XC", "c")
	s = strings.ReplaceAll(s, "CD", "d")
	s = strings.ReplaceAll(s, "CM", "m")

	sum := 0
	cs := strings.Split(s, "")
	for _, c := range cs {
		switch c {
		case "I":
			sum += 1
		case "V":
			sum += 5
		case "X":
			sum += 10
		case "L":
			sum += 50
		case "C":
			sum += 100
		case "D":
			sum += 500
		case "M":
			sum += 1000
		case "v":
			sum += 4
		case "x":
			sum += 9
		case "l":
			sum += 40
		case "c":
			sum += 90
		case "d":
			sum += 400
		case "m":
			sum += 900
		}
	}
	return sum
}

// @lc code=end
