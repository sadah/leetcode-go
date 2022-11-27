/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func isUgly(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	ret := PrimeFactors(n)
	return ret[len(ret)-1] <= 5
}

// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// @lc code=end
