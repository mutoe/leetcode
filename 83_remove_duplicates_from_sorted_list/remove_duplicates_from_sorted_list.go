package remove_duplicates_from_sorted_list

// https://leetcode.com/problems/remove-duplicates-from-sorted-list

// level: 1
// time: O(n) 4ms 83.90%
// space: O(n) 3.4M 7.63%

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
	m := make(map[int]bool)
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil {
		cur := prev.Next
		if m[cur.Val] {
			prev.Next = cur.Next
		} else {
			m[cur.Val] = true
			prev = prev.Next
		}
	}
	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
