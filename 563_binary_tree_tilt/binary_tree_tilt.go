package binary_tree_tilt

// https://leetcode.com/problems/binary-tree-tilt

// level: 1
// time: O(n) 8ms 94.08%
// space: O(1) 6.1M 39.25%

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
func findTilt(root *TreeNode) int {
	tilt(root)
	return sum(root)
}

func sum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return sum(node.Left) + sum(node.Right) + node.Val
}

func tilt(node *TreeNode) int {
	if node == nil {
		return 0
	}
	val := node.Val
	leftVal := tilt(node.Left)
	rightVal := tilt(node.Right)
	node.Val = diff(leftVal, rightVal)
	return leftVal + rightVal + val
}

func diff(a, b int) int {
	d := a - b
	if d >= 0 {
		return d
	}
	return -d
}

// leetcode submit region end(Prohibit modification and deletion)
