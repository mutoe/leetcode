package path_sum

// https://leetcode.com/problems/path-sum

// level: 1
// time: O(n) 4ms 96.52%
// space: O(1) 4.8M 14.98%

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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// leetcode submit region end(Prohibit modification and deletion)
