package binary_tree_inorder_traversal

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test1",
			args: []int{1, NULL, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "test2",
			args: []int{},
			want: []int{},
		},
		{
			name: "test3",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "test4",
			args: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "test5",
			args: []int{1, NULL, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args)
			got := inorderTraversal(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
