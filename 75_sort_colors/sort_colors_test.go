package sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
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
			args: args{[]int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test2",
			args: args{[]int{2, 0, 1}},
			want: []int{0, 1, 2},
		},
		{
			name: "test3",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "test4",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "test5",
			args: args{[]int{2, 1, 2, 0, 1, 1, 1, 1}},
			want: []int{0, 1, 1, 1, 1, 1, 2, 2},
		},
		{
			name: "test6",
			args: args{[]int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test7",
			args: args{[]int{0, 0}},
			want: []int{0, 0},
		},
		{
			name: "test8",
			args: args{[]int{1, 1}},
			want: []int{1, 1},
		},
		{
			name: "test9",
			args: args{[]int{2, 2}},
			want: []int{2, 2},
		},
		{
			name: "test10",
			args: args{[]int{1, 2, 0}},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)

			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
