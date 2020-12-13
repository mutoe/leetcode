package symmetric_tree

// https://leetcode.com/problems/symmetric-tree

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.9M 100%

// TODO: Solve it using iteratively

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

func symmetricTree(node *TreeNode) {
	node.Left, node.Right = node.Right, node.Left
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		symmetricTree(q)
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

// leetcode submit region end(Prohibit modification and deletion)
