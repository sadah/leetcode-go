/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

package code

// https://ikatakos.com/pot/programming_algorithm/graph_theory/lowest_common_ancestor
// https://www.wantedly.com/companies/wantedly/post_articles/340347
// https://qiita.com/maebaru/items/0fec7d2987c4a1efaa8a

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
//lint:ignore U1000 //
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

// @lc code=end
