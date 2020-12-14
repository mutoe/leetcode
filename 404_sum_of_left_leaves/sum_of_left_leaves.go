package sum_of_left_leaves

// https://leetcode.com/problems/sum-of-left-leaves

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.7M 87.21%

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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)
