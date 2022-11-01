/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := depth(root.Left)
	r := depth(root.Right)

	abs := l - r
	if l-r < 0 {
		abs = r - l
	}

	if abs > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left)
	r := depth(node.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

// @lc code=end
