package reverse_linked_list

import (
	"reflect"
	"testing"

	. "../utils/list_node"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ArrayToListNode(tt.args.head)
			got := ListNodeToArray(reverseList(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
