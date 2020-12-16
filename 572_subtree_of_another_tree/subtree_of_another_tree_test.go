package subtree_of_another_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		s []int
		t []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{3, 4, 5, 1, 2, NULL, NULL, NULL, NULL, 0}, []int{4, 1, 2}},
			want: false,
		},
		{
			name: "test3",
			args: args{[]int{1}, []int{0}},
			want: false,
		},
		{
			name: "test4",
			args: args{[]int{-1, -4, 8, -6, -2, 3, 9, NULL, -5, NULL, NULL, 0, 7}, []int{3, 0, 5848}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sN := ArrayToTreeNode(tt.args.s)
			tN := ArrayToTreeNode(tt.args.t)
			if got := isSubtree(sN, tN); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
