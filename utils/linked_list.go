package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, value := range arr {
		cur.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		cur = cur.Next
	}
	return head.Next
}

func LinkedListToArray(node *ListNode) []int {
	ret := make([]int, 0)
	cur := node
	for cur != nil {
		ret = append(ret, cur.Val)
		cur = cur.Next
	}
	return ret
}
