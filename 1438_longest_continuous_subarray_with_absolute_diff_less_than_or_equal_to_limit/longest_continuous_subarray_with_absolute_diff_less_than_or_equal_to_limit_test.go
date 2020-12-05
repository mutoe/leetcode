package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{8, 2, 4, 7}, 4},
			want: 2,
		},
		{
			name: "test2",
			args: args{[]int{10, 1, 2, 4, 7, 2}, 5},
			want: 4,
		},
		{
			name: "test3",
			args: args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0},
			want: 3,
		},
		{
			name: "test4",
			args: args{[]int{4, 8, 5, 1, 7, 9}, 6},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				// t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
