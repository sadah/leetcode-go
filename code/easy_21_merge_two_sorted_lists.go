/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

package code

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current1 := list1
	current2 := list2

	ret := &ListNode{-1, nil}
	pos := ret

	for current1 != nil && current2 != nil {
		if current1.Val <= current2.Val {
			pos.Next = current1
			current1 = current1.Next
		} else {
			pos.Next = current2
			current2 = current2.Next
		}
		pos = pos.Next
	}
	if current1 != nil {
		pos.Next = current1
	} else {
		pos.Next = current2
	}
	// printAllNode(ret)
	return ret.Next
}

// @lc code=end
