package matrix_cells_in_distance_order

import (
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{1, 2, 0, 0},
			want: []int{0, 1},
		},
		{
			name: "test2",
			args: args{2, 2, 0, 1},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "test3",
			args: args{2, 3, 1, 2},
			want: []int{0, 1, 1, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectLen := tt.args.C * tt.args.R
			got := allCellsDistOrder(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0)
			if len(got) != expectLen {
				t.Errorf("len(#{got}) != #{expectLen}")
				return
			}

			d := make([]int, 0)
			for i := 0; i < expectLen; i++ {
				d = append(d, manhattanDistance(got[i][0], got[i][1], tt.args.r0, tt.args.c0))
			}
			for i := 0; i < expectLen; i++ {
				if tt.want[i] != d[i] {
					t.Errorf("expect distance %v, got %v (%v)", tt.want, d, got)
					return
				}
			}
		})
	}
}
