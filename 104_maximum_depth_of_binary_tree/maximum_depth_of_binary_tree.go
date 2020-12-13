package maximum_depth_of_binary_tree

// https://leetcode.com/problems/maximum-depth-of-binary-tree

// level: 1
// time: O(n) 4ms 90.15%
// space: O(1) 4.5M 18.95%

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
	depth := 1
	prevOrder(root, 1, &depth)

	return depth
}

func prevOrder(node *TreeNode, level int, depth *int) {
	if node == nil {
		return
	}
	if level > *depth {
		*depth = level
	}
	prevOrder(node.Left, level+1, depth)
	prevOrder(node.Right, level+1, depth)
}

// leetcode submit region end(Prohibit modification and deletion)
