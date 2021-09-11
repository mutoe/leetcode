package find_k_closest_elements

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 4, 5}, 4, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 3, 4, 5}, 4, -1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test3",
			args: args{[]int{1, 2, 3, 4, 5}, 4, 6},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "test4",
			args: args{[]int{1, 2, 3, 4, 5}, 2, 4},
			want: []int{3, 4},
		},
		{
			name: "test5",
			args: args{[]int{1, 2, 3, 4, 5}, 3, 3},
			want: []int{2, 3, 4},
		},
		{
			name: "test6",
			args: args{[]int{1, 1, 1, 10, 10, 10}, 1, 9},
			want: []int{10},
		},
		{
			name: "test7",
			args: args{[]int{-2, -1, 1, 2, 3, 4, 5}, 7, 3},
			want: []int{-2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "test8",
			args: args{[]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4},
			want: []int{0, 1, 1, 1, 2, 3, 6, 7, 8},
		},
		{
			name: "test9",
			args: args{[]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 10}, 9, 9},
			want: []int{1, 1, 1, 2, 3, 6, 7, 8, 10},
		},
		{
			name: "test10",
			args: args{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5},
			want: []int{3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements(%v, %v, %v) = %v, want %v", tt.args.arr, tt.args.k, tt.args.x, got, tt.want)
			}
		})
	}
}
