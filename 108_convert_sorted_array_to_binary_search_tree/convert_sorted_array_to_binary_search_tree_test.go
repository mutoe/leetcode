package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	. "../utils/tree_node"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{-10, -3, 0, 5, 9}},
			want: []int{0, -3, 9, -10, NULL, 5},
		},
		{
			name: "test2",
			args: args{[]int{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.args.nums)
			if tt.want == nil {
				if root != nil {
					t.Errorf("expect nil, but got %v", TreeNodeToArray(root))
				}
				return
			}

			leftDepth := GetTreeDepth(root.Left)
			rightDepth := GetTreeDepth(root.Right)
			if leftDepth-rightDepth > 1 || leftDepth-rightDepth < -1 {
				t.Errorf("left is %v, right is %v", leftDepth, rightDepth)
			}
		})
	}
}
