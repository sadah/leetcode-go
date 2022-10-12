/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	var levels [][]int
	if root == nil {
		return levels
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		levels = append(levels, []int{})
		for levelSize := len(queue); levelSize > 0; levelSize-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			levels[level] = append(levels[level], queue[0].Val)
			queue = queue[1:]
		}
	}
	return levels
}

// @lc code=end
