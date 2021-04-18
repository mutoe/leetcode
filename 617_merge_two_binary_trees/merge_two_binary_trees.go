package merge_two_binary_trees

// https://leetcode.com/problems/merge-two-binary-trees

// level: 1
// time: O(n) 28ms 87.26%
// space: O(n) 8.6M 91.72%

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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	val := t1.Val + t2.Val
	left := mergeTrees(t1.Left, t2.Left)
	right := mergeTrees(t1.Right, t2.Right)
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

// leetcode submit region end(Prohibit modification and deletion)
