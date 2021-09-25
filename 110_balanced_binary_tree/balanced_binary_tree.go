package balanced_binary_tree

// https://leetcode.com/problems/balanced-binary-tree

// level: 1
// time: O(n) 4ms 96.71%
// space: O(n) 5.9M 40.79%

import (
	. "../utils/tree_node"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func helper(node *TreeNode) (float64, bool) {
	if node == nil {
		return 0, true
	}
	left, bl := helper(node.Left)
	right, br := helper(node.Right)
	if math.Abs(left-right) > 1 {
		return -1, false
	}
	return math.Max(left, right) + 1, bl && br
}

func isBalanced(root *TreeNode) bool {
	_, ret := helper(root)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
