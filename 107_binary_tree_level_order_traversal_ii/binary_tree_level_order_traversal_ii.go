package binary_tree_level_order_traversal_ii

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.8M 99.24%

// TODO: Need to review

import (
	. "../utils/tree_node"
)

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		arr := make([]int, 0)
		for _, node := range queue {
			queue = queue[1:]
			if node == nil {
				continue
			}
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, arr)
	}
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
