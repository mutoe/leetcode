package lowest_common_ancestor_of_a_binary_search_tree

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root []int
		p    []int
		q    []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
				p:    []int{2, 0, 4, NULL, NULL, 3, 5},
				q:    []int{8, 7, 9},
			},
			want: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			q := ArrayToTreeNode(tt.args.q)
			p := ArrayToTreeNode(tt.args.p)
			got := TreeNodeToArray(lowestCommonAncestor(root, p, q))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
