package binary_tree_tilt

import (
	"testing"

	. "../utils/tree_node"
)

func Test_findTilt(t *testing.T) {
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
			args: args{[]int{1, 2, 3}},
			want: 1,
		},
		{
			name: "test2",
			args: args{[]int{4, 2, 9, 3, 5, NULL, 7}},
			want: 15,
		},
		{
			name: "test3",
			args: args{[]int{21, 7, 14, 1, 1, 2, 2, 3, 3}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := findTilt(root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
