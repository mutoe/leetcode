package search_in_a_binary_search_tree

// https://leetcode.com/problems/search-in-a-binary-search-tree

// level: 1
// time: O(n) 24ms 84.67%
// space: O(1) 6.6M 57.33%

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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return nil
}

// leetcode submit region end(Prohibit modification and deletion)
