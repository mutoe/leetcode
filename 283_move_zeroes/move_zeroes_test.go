package move_zeroes

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
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
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "test2",
			args: args{[]int{0, 0, 1}},
			want: []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// original := append([]int(nil), tt.args.nums...)
			moveZeroes(tt.args.nums)
			for i := 0; i < len(tt.want); i++ {
				if tt.want[i] != tt.args.nums[i] {
					t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
					return
				}
			}
		})
	}
}
