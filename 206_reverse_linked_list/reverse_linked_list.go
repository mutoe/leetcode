package reverse_linked_list

// https://leetcode.com/problems/reverse-linked-list

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.7M 36.88%

// TODO: using recursive call to resolve

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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	for ; head != nil; head = head.Next {
		dummyHead.Next = &ListNode{
			Val:  head.Val,
			Next: dummyHead.Next,
		}
	}
	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
