package middle_of_the_linked_list

import (
	. "../utils/list_node"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
			want: []int{3, 4, 5},
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 3, 4, 5, 6}},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToArray(middleNode(ArrayToListNode(tt.args.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
