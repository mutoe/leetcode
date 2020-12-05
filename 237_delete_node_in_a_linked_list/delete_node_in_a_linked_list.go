package delete_node_in_a_linked_list

// https://leetcode.com/problems/delete-node-in-a-linked-list

// level: 1
// time: O(1) 0ms 100%
// space: O(1) 2.9M 99.66%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}

// leetcode submit region end(Prohibit modification and deletion)
