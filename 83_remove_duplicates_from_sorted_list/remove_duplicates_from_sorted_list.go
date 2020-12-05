package remove_duplicates_from_sorted_list

// https://leetcode.com/problems/remove-duplicates-from-sorted-list

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 3.2M 16.53%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	for prev.Next != nil {
		if prev.Val == prev.Next.Val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

// leetcode submit region end(Prohibit modification and deletion)
