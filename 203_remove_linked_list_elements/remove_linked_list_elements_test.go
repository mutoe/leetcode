package remove_linked_list_elements

import (
	"reflect"
	"testing"

	. "../utils/list_node"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args{
				head: []int{},
				val:  1,
			},
			want: []int{},
		},
		{
			name: "test3",
			args: args{
				head: []int{1, 1},
				val:  1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ArrayToLinkedList(tt.args.head)
			retArr := LinkedListToArray(removeElements(head, tt.args.val))
			if !reflect.DeepEqual(retArr, tt.want) {
				t.Errorf("removeElements() = %v, want %v", retArr, tt.want)
			}
		})
	}
}
