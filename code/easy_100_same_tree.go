/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */
package code

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//lint:ignore U1000 //
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

// @lc code=end
