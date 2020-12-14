package find_mode_in_binary_search_tree

import (
	"reflect"
	"sort"
	"testing"

	. "../utils/tree_node"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, NULL, 2, 2}},
			want: []int{2},
		},
		{
			name: "test2",
			args: args{[]int{1, NULL, 1, NULL, 2, 2}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			got := findMode(root)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
