package binary_tree_level_order_traversal

// https://leetcode.com/problems/binary-tree-level-order-traversal

// level: 2
// time: O(n) 0ms 100%
// space: O(n) 2.8M 100%

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var ret [][]int
	for len(queue) > 0 {
		var arr []int
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
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
