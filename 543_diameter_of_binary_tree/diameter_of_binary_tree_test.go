package diameter_of_binary_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 3,
		},
		{
			name: "test2",
			args: args{[]int{1}},
			want: 0,
		},
		{
			name: "test3",
			args: args{[]int{1, 2, NULL, 4, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := diameterOfBinaryTree(root); got != tt.want {
				// t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
