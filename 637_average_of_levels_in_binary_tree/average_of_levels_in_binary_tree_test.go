package average_of_levels_in_binary_tree

import (
	"reflect"
	"testing"

	. "../utils/tree_node"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test1",
			args: args{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			want: []float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ArrayToTreeNode(tt.args.root)
			if got := averageOfLevels(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
