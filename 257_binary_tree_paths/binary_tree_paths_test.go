package binary_tree_paths

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, NULL, 5}},
			want: []string{"1->2->5", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := binaryTreePaths(root); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
