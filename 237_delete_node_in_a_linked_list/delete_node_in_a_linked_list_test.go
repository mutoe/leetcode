package delete_node_in_a_linked_list

import (
	"reflect"
	"testing"
)
import . "../utils/list_node"

func Test_deleteNode(t *testing.T) {
	type args struct {
		head []int
		node int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				head: []int{4, 5, 1, 9},
				node: 5,
			},
			want: []int{4, 1, 9},
		},
		{
			name: "test2",
			args: args{
				head: []int{4, 5, 1, 9},
				node: 1,
			},
			want: []int{4, 5, 9},
		},
		{
			name: "test3",
			args: args{
				head: []int{1, 2, 3, 4},
				node: 3,
			},
			want: []int{1, 2, 4},
		},
		{
			name: "test4",
			args: args{
				head: []int{0, 1},
				node: 0,
			},
			want: []int{1},
		},
		{
			name: "test5",
			args: args{
				head: []int{-3, 5, -99},
				node: -3,
			},
			want: []int{5, -99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ArrayToListNode(tt.args.head)
			node := head
			for node != nil {
				if node.Val == tt.args.node {
					break
				}
				node = node.Next
			}

			deleteNode(node)

			if got := ListNodeToArray(head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
