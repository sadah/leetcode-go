/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	l := root.Left
	r := root.Right
	root.Left = r
	root.Right = l
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

// @lc code=end
