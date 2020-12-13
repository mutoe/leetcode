package convert_sorted_array_to_binary_search_tree

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree

// level: 1
// time: O(n) 108ms 62.90%
// space: O(n) 67.4M 97.53%

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	r := len(nums) - 1
	mid := r >> 1
	root := &TreeNode{Val: nums[mid]}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[0:mid])
	}
	if r > mid {
		root.Right = sortedArrayToBST(nums[mid+1 : r+1])
	}
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
