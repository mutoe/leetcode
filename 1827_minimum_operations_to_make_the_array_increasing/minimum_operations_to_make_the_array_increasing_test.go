package minimum_operations_to_make_the_array_increasing

import "testing"

func Test_minOperations(t *testing.T) {
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
			args: args{[]int{1, 1, 1}},
			want: 3,
		},
		{
			name: "test2",
			args: args{[]int{1, 5, 2, 4, 1}},
			want: 14,
		},
		{
			name: "test3",
			args: args{[]int{8}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
