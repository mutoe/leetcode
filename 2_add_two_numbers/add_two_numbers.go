package add_two_numbers

// https://leetcode.com/problems/add-two-numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n) 12ms 64.05%
// space: O(n) 4.9M

// leetcode submit region begin(Prohibit modification and deletion)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := ListNode{}
	p := &root
	for {
		if l1 == nil {
			p.Next = l2
			break
		} else if l2 == nil {
			p.Next = l1
			break
		} else {
			p.Next = &ListNode{}
			p.Next.Val = l1.Val + l2.Val
			p = p.Next
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
	}
	p = &root
	for p != nil {
		if p.Val >= 10 {
			if p.Next != nil {
				p.Next.Val++
			} else {
				p.Next = &ListNode{Val: 1}
			}
			p.Val -= 10
		}
		p = p.Next
	}
	return root.Next
}

// leetcode submit region end(Prohibit modification and deletion)
