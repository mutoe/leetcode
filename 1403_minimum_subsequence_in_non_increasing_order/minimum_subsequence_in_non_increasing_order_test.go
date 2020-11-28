package minimum_subsequence_in_non_increasing_order

import (
	"reflect"
	"testing"
)

func Test_minSubsequence(t *testing.T) {
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
			args: args{[]int{4, 3, 10, 9, 8}},
			want: []int{10, 9},
		},
		{
			name: "test2",
			args: args{[]int{4, 4, 7, 6, 7}},
			want: []int{7, 7, 6},
		},
		{
			name: "test3",
			args: args{[]int{6}},
			want: []int{6},
		},
		{
			name: "test4",
			args: args{[]int{10, 2, 5}},
			want: []int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubsequence(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
