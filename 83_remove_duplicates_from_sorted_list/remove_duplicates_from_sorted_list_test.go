package remove_duplicates_from_sorted_list

import (
	. "../utils/list_node"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args: args{[]int{1, 1, 2}},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{[]int{1, 1, 2, 3, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToArray(deleteDuplicates(ArrayToListNode(tt.args.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
