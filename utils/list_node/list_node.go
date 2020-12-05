package list_node

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

func ArrayToLinkedListReturnTail(arr []int) (head, tail *ListNode) {
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, value := range arr {
		cur.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		cur = cur.Next
		tail = cur
	}
	head = dummyHead.Next
	return
}

func ArrayToLinkedListWithTailCycle(arr []int, pos int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	var tail *ListNode
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		cur.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		if i == pos {
			tail = cur.Next
		}
		if i == len(arr)-1 {
			cur.Next.Next = tail
		}
		cur = cur.Next
	}
	return dummyHead.Next
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
