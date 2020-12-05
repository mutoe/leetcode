package remove_linked_list_elements

// https://leetcode.com/problems/remove-linked-list-elements

// level: 1
// time: O(n) 8ms 88.89%
// space: O(1) 4.7M 100%

import (
	. "../utils/list_node"
)

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil {
		next := cur.Next
		if next != nil && next.Val == val {
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
