package find_mode_in_binary_search_tree

// https://leetcode.com/problems/find-mode-in-binary-search-tree

// level: 1
// time: O(n) 12ms 51.16%
// space: O(n) 6.8M 6.98%

// TODO: Optimize space (without using any extra space)

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
func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		m[node.Val]++
		q = append(q, node.Left, node.Right)
	}
	var ret []int
	max := 0
	for v, count := range m {
		if count > max {
			max = count
			ret = []int{v}
		} else if count == max {
			ret = append(ret, v)
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
