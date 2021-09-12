package number_of_pairs_of_interchangeable_rectangles

import "testing"

func Test_interchangeableRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{[][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}},
			want: 6,
		},
		{
			name: "test2",
			args: args{[][]int{{4, 5}, {7, 8}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interchangeableRectangles(tt.args.rectangles); got != tt.want {
				t.Errorf("interchangeableRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
