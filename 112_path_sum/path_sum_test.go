package path_sum

import (
	"testing"

	. "../utils/tree_node"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1}, 22},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{-2, NULL, -3}, -5},
			want: true,
		},
		{
			name: "test3",
			args: args{[]int{1, -2, -3, 1, 3, -2, NULL, -1}, -1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := hasPathSum(root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
