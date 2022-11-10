/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	return potl(root, []int{})
}

func potl(root *TreeNode, ret []int) []int {
	if root != nil {
		ret = potl(root.Left, ret)
		ret = potl(root.Right, ret)
		ret = append(ret, root.Val)
	}
	return ret
}

// @lc code=end
