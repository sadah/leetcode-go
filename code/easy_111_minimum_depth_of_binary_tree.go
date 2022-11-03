/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */

package code

import "math"

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := math.MaxInt64
	if root.Left != nil {
		min = int(math.Min(float64(min), float64(minDepth(root.Left))))
	}
	if root.Right != nil {
		min = int(math.Min(float64(min), float64(minDepth(root.Right))))
	}
	return min + 1
}

// @lc code=end
