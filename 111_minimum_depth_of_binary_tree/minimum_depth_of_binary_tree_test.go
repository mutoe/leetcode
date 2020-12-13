package minimum_depth_of_binary_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{[]int{3, 9, 20, NULL, NULL, 15, 7}},
		// 	want: 2,
		// },
		// {
		// 	name: "test2",
		// 	args: args{[]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6}},
		// 	want: 5,
		// },
		// {
		// 	name: "test3",
		// 	args: args{[]int{1,2,3,4,5}},
		// 	want: 2,
		// },
		{
			name: "test4",
			args: args{[]int{0, 2, 4, 1, NULL, 3, -1, 5, 1, NULL, 6, NULL, 8}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(ArrayToTreeNode(tt.args.root)); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
