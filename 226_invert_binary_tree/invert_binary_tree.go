package invert_binary_tree

// https://leetcode.com/problems/invert-binary-tree

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.2M 100%

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
