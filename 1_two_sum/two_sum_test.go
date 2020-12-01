package two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{0, 1},
		},
		{
			name: "test2",
			args: args{[]int{3, 2, 4}, 6},
			want: []int{1, 2},
		},
		{
			name: "test3",
			args: args{[]int{3, 2, 4}, 4},
			want: []int{},
		},
		{
			name: "test4",
			args: args{[]int{0, 4, 3, 0}, 0},
			want: []int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
