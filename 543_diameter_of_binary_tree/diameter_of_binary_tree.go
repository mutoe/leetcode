package diameter_of_binary_tree

// https://leetcode.com/problems/diameter-of-binary-tree

// level: 1

import . "../utils/tree_node"

// TODO: Unresolved problem

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return diameterOfBinaryTree(root.Left) + diameterOfBinaryTree(root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)
