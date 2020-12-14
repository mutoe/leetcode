package invert_binary_tree

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_invertTree(t *testing.T) {
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
			args: args{[]int{4, 2, 7, 1, 3, 6, 9}},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			got := TreeNodeToArray(invertTree(root))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
