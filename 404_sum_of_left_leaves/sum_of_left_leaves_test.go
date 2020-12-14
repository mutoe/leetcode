package sum_of_left_leaves

import (
	"testing"

	. "../utils/tree_node"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
			args: args{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := sumOfLeftLeaves(root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
