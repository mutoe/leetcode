package count_negative_numbers_in_a_sorted_matrix

import "testing"

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			}},
			want: 8,
		},
		{
			name: "test2",
			args: args{[][]int{{3, 2}, {1, 0}}},
			want: 0,
		},
		{
			name: "test3",
			args: args{[][]int{{1, -1}, {-1, -1}}},
			want: 3,
		},
		{
			name: "test4",
			args: args{[][]int{{-1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
