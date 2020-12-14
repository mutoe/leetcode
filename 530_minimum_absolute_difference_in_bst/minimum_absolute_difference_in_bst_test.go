package minimum_absolute_difference_in_bst

import (
	"testing"

	. "../utils/tree_node"
)

func Test_getMinimumDifference(t *testing.T) {
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
			args: args{[]int{1, NULL, 3, 2}},
			want: 1,
		},
		{
			name: "test2",
			args: args{[]int{236, 104, 701, NULL, 227, NULL, 911}},
			want: 9,
		},
		{
			name: "test3",
			args: args{[]int{1, NULL, 5, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := getMinimumDifference(root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
