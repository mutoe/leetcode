package middle_of_the_linked_list

// https://leetcode.com/problems/middle-of-the-linked-list

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M 100%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// leetcode submit region end(Prohibit modification and deletion)
