package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{[]int{3, 9, 20, 0, 0, 15, 7}},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := levelOrder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
