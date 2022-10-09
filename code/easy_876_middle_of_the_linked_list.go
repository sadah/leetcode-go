/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */

package code

// @lc code=start
//
// Fast and Slow Pointer
//
//lint:ignore U1000 //
func middleNode(head *ListNode) *ListNode {
	mid := head
	pos := head
	for pos != nil && pos.Next != nil {
		mid = mid.Next
		pos = pos.Next.Next
	}
	return mid
}

//lint:ignore U1000 //
func middleNodeByIndex(head *ListNode) *ListNode {
	len := 0
	pos := head
	for pos != nil {
		len++
		pos = pos.Next
	}
	mid := int(len / 2)
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}

// @lc code=end
