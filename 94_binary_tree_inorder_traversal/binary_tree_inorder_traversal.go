package binary_tree_inorder_traversal

// https://leetcode.com/problems/binary-tree-inorder-traversal

// level: 1
// time: O(N) 0ms 100%
// space: 2.1M 99.77%

import . "../utils/tree_node"

//leetcode submit region begin(Prohibit modification and deletion)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
