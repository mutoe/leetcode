package maximum_depth_of_binary_tree

// https://leetcode.com/problems/maximum-depth-of-binary-tree

// level: 1
// time: O(n) 4ms 90.15%
// space: O(1) 4.4M 18.95%

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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// leetcode submit region end(Prohibit modification and deletion)
