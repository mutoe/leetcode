package lowest_common_ancestor_of_a_binary_search_tree

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree

// level: 1
// time: O(n) 20ms 80.11%
// space: O(1) 6.8M 5.52%

import . "../utils/tree_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
