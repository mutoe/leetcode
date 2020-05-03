package check_if_all_1s_are_at_least_length_k_places_away

import "testing"

func Test_kLengthApart(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 0, 0, 1, 0, 1}, 2},
			want: false,
		},
		{
			name: "test3",
			args: args{[]int{1, 1, 1, 1, 1}, 0},
			want: true,
		},
		{
			name: "test4",
			args: args{[]int{0, 1, 0, 1}, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLengthApart(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
