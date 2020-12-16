package average_of_levels_in_binary_tree

// https://leetcode.com/problems/average-of-levels-in-binary-tree

// level: 1
// time: O(n) 8ms 89.47%
// space: O(n) 6.1M 89.47%

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
func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	var ret []float64
	for len(q) > 0 {
		sum := 0
		count := 0
		for _, node := range q {
			q = q[1:]
			if node == nil {
				continue
			}
			count++
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}
		if count > 0 {
			ret = append(ret, float64(sum)/float64(count))
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
