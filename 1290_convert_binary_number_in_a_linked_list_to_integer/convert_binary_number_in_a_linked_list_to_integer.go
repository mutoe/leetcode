package convert_binary_number_in_a_linked_list_to_integer

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.1M 5.83%

import . "../utils/list_node"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var num int
	for head != nil {
		num = num << 1
		num += head.Val
		head = head.Next
	}
	return num
}

// leetcode submit region end(Prohibit modification and deletion)
