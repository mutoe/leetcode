package count_number_of_pairs_with_absolute_difference_k

import "testing"

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 2, 1}, 1},
			want: 4,
		},
		{
			name: "test2",
			args: args{[]int{1, 3}, 3},
			want: 0,
		},
		{
			name: "test3",
			args: args{[]int{3, 2, 1, 5, 4}, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
