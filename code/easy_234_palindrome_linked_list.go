/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

package code

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//lint:ignore U1000 //
func isPalindrome(head *ListNode) bool {
	l := []int{}
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	n := len(l)
	for i := 0; i < n/2; i++ {
		if l[i] != l[n-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end
