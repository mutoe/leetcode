package balanced_binary_tree

import "testing"
import . "../utils/tree_node"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 2, 3, 3, NULL, NULL, 4, 4}},
			want: false,
		},
		{
			name: "test3",
			args: args{[]int{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			got := isBalanced(root)
			if got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
