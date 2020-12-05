package linked_list_cycle

// https://leetcode.com/problems/linked-list-cycle

// level: 1
// time: O(n) 8ms 22.44%
// space: O(n) 5.2M 16.22%

import (
	. "../utils/list_node"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if m[cur] {
			return true
		}
		m[cur] = true
		cur = cur.Next
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
