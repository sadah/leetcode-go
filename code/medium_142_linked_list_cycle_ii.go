/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for head != nil {
		_, ok := m[head]
		if ok {
			return head
		} else {
			m[head] = true
			head = head.Next
		}
	}
	return nil
}

// @lc code=end
