package intersection_of_two_linked_lists

// https://leetcode.com/problems/intersection-of-two-linked-lists

// level: 1
// time: O(n) 44ms 30.10%
// space: O(n) 7.3M 67.82%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// leetcode submit region end(Prohibit modification and deletion)
