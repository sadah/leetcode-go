/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */

package code

import (
	"strconv"
	"strings"
)

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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return walkBTree(root, []string{strconv.Itoa(root.Val)}, []string{})
}

func walkBTree(root *TreeNode, strs []string, ret []string) []string {
	if root.Left == nil && root.Right == nil {
		s := strings.Join(strs, "->")
		return append(ret, s)
	}
	if root.Left != nil {
		s := append(strs, strconv.Itoa(root.Left.Val))
		ret = walkBTree(root.Left, s, ret)
	}
	if root.Right != nil {
		s := append(strs, strconv.Itoa(root.Right.Val))
		ret = walkBTree(root.Right, s, ret)
	}
	return ret
}

// @lc code=end
