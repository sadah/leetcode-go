/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

package code

// @lc code=start
//
//lint:ignore U1000 //
func preorder(root *Node) []int {
	var ret []int
	if root != nil {
		ret = walk(root, ret)
	}
	return ret
}

func walk(node *Node, ret []int) []int {
	ret = append(ret, node.Val)
	for _, n := range node.Children {
		ret = walk(n, ret)
	}
	return ret
}

// @lc code=end
