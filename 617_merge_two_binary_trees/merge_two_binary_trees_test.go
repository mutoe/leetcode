package merge_two_binary_trees

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 []int
		t2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, 3, 2, 5}, []int{2, 1, 3, NULL, 4, NULL, 7}},
			want: []int{3, 4, 5, 5, 4, NULL, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1 := ArrayToTreeNode(tt.args.t1)
			t2 := ArrayToTreeNode(tt.args.t2)
			got := TreeNodeToArray(mergeTrees(t1, t2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
