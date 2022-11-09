/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	return pot(root, []int{})
}

func pot(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	if root.Left != nil {
		ret = pot(root.Left, ret)
	}
	if root.Right != nil {
		ret = pot(root.Right, ret)
	}
	return ret
}

// @lc code=end
