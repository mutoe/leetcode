package merge_two_sorted_lists

import (
	"reflect"
	"testing"

	. "../utils/list_node"
)

func Test_mergeTwoLists(t *testing.T) {
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
				l1: []int{1, 2, 4},
				l2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "test2",
			args: args{
				l1: []int{},
				l2: []int{},
			},
			want: []int{},
		},
		{
			name: "test3",
			args: args{
				l1: []int{},
				l2: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1, l2 := ArrayToListNode(tt.args.l1), ArrayToListNode(tt.args.l2)
			if got := ListNodeToArray(mergeTwoLists(l1, l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
