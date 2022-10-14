/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

package code

import (
	"fmt"
	"math"
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
func isValidBST(root *TreeNode) bool {
	prev = -math.MaxInt64
	return InorderTraversal(root)
}

var prev int

func InorderTraversal(node *TreeNode) bool {
	if node == nil {
		return true
	}
	fmt.Println("1: node.Val: ", node.Val, " prev: ", prev)

	if !InorderTraversal(node.Left) {
		return false
	}
	if node.Val <= prev {
		return false
	}
	//fmt.Println(node.Val)
	prev = node.Val
	return InorderTraversal(node.Right)
}

func PreorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	if node.Left != nil {
		PreorderTraversal(node.Left)
	}
	if node.Right != nil {
		PreorderTraversal(node.Right)
	}
}

// @lc code=end
