package linked_list_cycle

// https://leetcode.com/problems/linked-list-cycle

// level: 1
// time: O(n) 4ms 96%
// space: O(1) 3.8M 24.67%

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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for slow, fast := head, head.Next; slow != fast; slow, fast = slow.Next, fast.Next.Next {
		if slow == nil || fast == nil || fast.Next == nil {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
