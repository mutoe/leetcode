package subtree_of_another_tree

// https://leetcode.com/problems/subtree-of-another-tree

// level: 1
// time: O(n^2) 20ms 47.97%
// space: O(1) 6.2M 34.15%

import . "../utils/tree_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSameTree(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s != nil && t != nil {
		return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
