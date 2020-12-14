package minimum_absolute_difference_in_bst

// https://leetcode.com/problems/minimum-absolute-difference-in-bst

// level: 1
// time: O(n) 12ms 72.92%
// space: O(1) 6M 41.67%

import (
	"math"

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
func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt32
	prev := -1
	traverse(root, &prev, &min)
	return min
}

func traverse(node *TreeNode, prev, min *int) {
	if node == nil {
		return
	}
	traverse(node.Left, prev, min)
	if *prev == -1 {
		*prev = node.Val
	} else if node.Val-*prev < *min {
		*min = node.Val - *prev
	}
	*prev = node.Val
	traverse(node.Right, prev, min)
}

// leetcode submit region end(Prohibit modification and deletion)
