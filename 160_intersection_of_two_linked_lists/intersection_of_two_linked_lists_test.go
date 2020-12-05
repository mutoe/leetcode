package intersection_of_two_linked_lists

import (
	"testing"

	. "../utils/list_node"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		intersectVal int
		listA        []int
		listB        []int
		skipA        int
		skipB        int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				intersectVal: 8,
				listA:        []int{4, 1, 8, 4, 5},
				listB:        []int{5, 6, 1, 8, 4, 5},
				skipA:        2,
				skipB:        3,
			},
		},
		{
			name: "test2",
			args: args{
				intersectVal: 2,
				listA:        []int{1, 9, 1, 2, 4},
				listB:        []int{3, 2, 4},
				skipA:        3,
				skipB:        1,
			},
		},
		{
			name: "test3",
			args: args{
				intersectVal: 0,
				listA:        []int{2, 6, 4},
				listB:        []int{1, 5},
				skipA:        3,
				skipB:        2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, intersection := generateIntersectionLinkedLists(tt.args.intersectVal, tt.args.listA, tt.args.listB, tt.args.skipA, tt.args.skipB)
			ret := getIntersectionNode(headA, headB)
			if ret != intersection {
				retArr := ListNodeToArray(ret)
				expectArr := ListNodeToArray(intersection)
				t.Errorf("getIntersectionNode() = %v, want %v", retArr, expectArr)
			}
		})
	}
}

func generateIntersectionLinkedLists(intersectionVal int, listA, listB []int, skipA, skipB int) (headA, headB, intersection *ListNode) {
	if intersectionVal == 0 {
		headA = ArrayToListNode(listA)
		headB = ArrayToListNode(listB)
		intersection = nil
		return
	}

	var tailA, tailB *ListNode
	headA, tailA = ArrayToListNodeReturnTail(listA[:skipA])
	headB, tailB = ArrayToListNodeReturnTail(listB[:skipB])
	intersection = ArrayToListNode(listA[skipA:])
	tailA.Next = intersection
	tailB.Next = intersection
	return
}
