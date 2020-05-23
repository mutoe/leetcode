package add_two_numbers

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				[]int{2, 4, 3},
				[]int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "test2",
			args: args{
				[]int{1, 8},
				[]int{0},
			},
			want: []int{1, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := ListNodeToArray(addTwoNumbers(ArrayToListNode(tt.args.l1), ArrayToListNode(tt.args.l2)))
			if !compareArray(actualValue, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", actualValue, tt.want)
			}
		})
	}
}

func ArrayToListNode(array []int) *ListNode {
	root := ListNode{}
	p := &root
	for _, val := range array {
		p.Next = &ListNode{}
		p.Next.Val = val
		p = p.Next
	}
	return root.Next
}

func ListNodeToArray(list *ListNode) (ret []int) {
	for p := list; p != nil; p = p.Next {
		ret = append(ret, p.Val)
	}
	return
}

func compareArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
