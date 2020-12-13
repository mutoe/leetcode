package maximum_depth_of_binary_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_maxDepth(t *testing.T) {
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
			args: args{[]int{3, 9, 20, 0, 0, 15, 7}},
			want: 3,
		},
		{
			name: "test2",
			args: args{[]int{1, 0, 2}},
			want: 2,
		},
		{
			name: "test3",
			args: args{[]int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
