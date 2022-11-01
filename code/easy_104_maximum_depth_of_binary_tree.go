/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxNodeDepth(root, 0)
}

func maxNodeDepth(node *TreeNode, d int) int {
	if node == nil {
		return d
	}
	l := maxNodeDepth(node.Left, d+1)
	r := maxNodeDepth(node.Right, d+1)
	if l > r {
		return l
	} else {
		return r
	}
}

// @lc code=end
