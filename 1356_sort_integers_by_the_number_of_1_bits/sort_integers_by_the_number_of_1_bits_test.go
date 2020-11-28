package sort_integers_by_the_number_of_1_bits

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			name: "test2",
			args: args{
				arr: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			},
			want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			name: "test3",
			args: args{
				arr: []int{10000, 10000},
			},
			want: []int{10000, 10000},
		},
		{
			name: "test4",
			args: args{
				arr: []int{2, 3, 5, 7, 11, 13, 17, 19},
			},
			want: []int{2, 3, 5, 17, 7, 11, 13, 19},
		},
		{
			name: "test5",
			args: args{
				arr: []int{10, 100, 1000, 10000},
			},
			want: []int{10, 100, 10000, 1000},
		},
		{
			name: "test6",
			args: args{
				arr: []int{10000},
			},
			want: []int{10000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
