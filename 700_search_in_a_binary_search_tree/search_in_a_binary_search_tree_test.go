package search_in_a_binary_search_tree

import (
	"reflect"
	"testing"
)
import . "../utils/tree_node"

func Test_searchBST(t *testing.T) {
	type args struct {
		root []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{4, 2, 7, 1, 3}, 2},
			want: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			got := TreeNodeToArray(searchBST(root, tt.args.val))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
