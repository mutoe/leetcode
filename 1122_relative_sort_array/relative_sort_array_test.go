package relative_sort_array

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
				arr2: []int{2, 1, 4, 3, 9, 6},
			},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "test2",
			args: args{
				arr1: []int{28, 6, 22, 8, 44, 17},
				arr2: []int{22, 28, 8, 6},
			},
			want: []int{22, 28, 8, 6, 17, 44},
		},
		{
			name: "test3",
			args: args{
				arr1: []int{26, 21, 11, 20, 50, 34, 1, 18},
				arr2: []int{21, 11, 26, 20},
			},
			want: []int{21, 11, 26, 20, 1, 18, 34, 50},
		},
		{
			name: "test4",
			args: args{
				arr1: []int{33, 22, 48, 4, 39, 36, 41, 47, 15, 45},
				arr2: []int{22, 33, 48, 4},
			},
			want: []int{22, 33, 48, 4, 15, 36, 39, 41, 45, 47},
		},
		{
			name: "test5",
			args: args{
				arr1: []int{943, 790, 427, 722, 860, 550, 225, 846, 715, 320},
				arr2: []int{943, 715, 427, 790, 860, 722, 225, 320, 846, 550},
			},
			want: []int{943, 715, 427, 790, 860, 722, 225, 320, 846, 550},
		},
		{
			name: "test6",
			args: args{
				arr1: []int{1, 2, 3, 4},
				arr2: []int{2, 1, 3, 4},
			},
			want: []int{2, 1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
