package minimum_depth_of_binary_tree

// https://leetcode.com/problems/minimum-depth-of-binary-tree

// level: 1
// time: O(n) 228ms 85.67%
// space: O(n) 21.2M 19.28%

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
func minDepth(root *TreeNode) int {
	depth := 1
	for root != nil {
		if root.Left == nil && root.Right == nil {
			return depth
		} else if root.Right == nil {
			depth++
			root = root.Left
		} else if root.Left == nil {
			depth++
			root = root.Right
		} else {
			break
		}
	}

	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left < right {
		return left + depth
	} else {
		return right + depth
	}
}

// leetcode submit region end(Prohibit modification and deletion)
