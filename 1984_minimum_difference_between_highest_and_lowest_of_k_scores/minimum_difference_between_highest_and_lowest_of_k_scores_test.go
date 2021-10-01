package minimum_difference_between_highest_and_lowest_of_k_scores

import "testing"

func Test_minimumDifference(t *testing.T) {
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
			args: args{[]int{90}, 1},
			want: 0,
		},
		{
			name: "test2",
			args: args{[]int{9, 4, 1, 7}, 2},
			want: 2,
		},
		{
			name: "test3",
			args: args{[]int{9, 4, 1, 7}, 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
