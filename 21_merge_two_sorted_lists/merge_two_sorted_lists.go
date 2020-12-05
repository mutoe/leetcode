package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.5M 48.96%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{Next: l1}
	prev := dummy

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			prev.Next = l2
			l2 = l2.Next
			prev = prev.Next
			prev.Next = l1
		} else {
			prev, l1 = l1, l1.Next
		}
	}
	if l1 == nil {
		prev.Next = l2
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
