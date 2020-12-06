package check_array_formation_through_concatenation

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{85}, [][]int{{85}}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{15, 88}, [][]int{{88}, {15}}},
			want: true,
		},
		{
			name: "test3",
			args: args{[]int{49, 18, 16}, [][]int{{16, 18, 49}}},
			want: false,
		},
		{
			name: "test4",
			args: args{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}},
			want: true,
		},
		{
			name: "test5",
			args: args{[]int{1, 3, 5, 7}, [][]int{{2, 4, 6, 8}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
