/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func reverseList(head *ListNode) *ListNode {
	var ret *ListNode = nil

	for head != nil {
		tmpNext := head.Next
		head.Next = ret
		ret = head
		head = tmpNext
		// PrintAllNode(ret)
	}
	return ret
}

// @lc code=end
