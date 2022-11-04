/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	sum := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// @lc code=end
