package same_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p []int
		q []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3}, []int{1, 2, 3}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2}, []int{1, 0, 2}},
			want: false,
		},
		{
			name: "test3",
			args: args{[]int{1, 2, 1}, []int{1, 1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ArrayToTreeNode(tt.args.p)
			q := ArrayToTreeNode(tt.args.q)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
