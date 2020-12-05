package reverse_linked_list

// https://leetcode.com/problems/reverse-linked-list

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.5M 71.38%

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

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	cur := head
//
// 	for cur != nil {
// 		cur.Next, prev, cur = prev, cur, cur.Next
// 	}
//
// 	return prev
// }

// Recursive version
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
