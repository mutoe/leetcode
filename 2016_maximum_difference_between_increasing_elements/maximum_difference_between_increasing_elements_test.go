package maximum_difference_between_increasing_elements

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{7, 1, 5, 4}},
			want: 4,
		},
		{
			name: "test2",
			args: args{[]int{9, 4, 3, 2}},
			want: -1,
		},
		{
			name: "test3",
			args: args{[]int{1, 5, 2, 10}},
			want: 9,
		},
		{
			name: "test4",
			args: args{[]int{1, 1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
