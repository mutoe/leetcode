package symmetric_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_isSymmetric(t *testing.T) {
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
			args: args{[]int{1, 2, 2, 3, 4, 4, 3}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 2, 0, 3, 0, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
