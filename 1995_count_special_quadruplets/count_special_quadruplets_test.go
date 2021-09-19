package count_special_quadruplets

import "testing"

func Test_countQuadruplets(t *testing.T) {
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
			args: args{[]int{1, 2, 3, 6}},
			want: 1,
		},
		{
			name: "test2",
			args: args{[]int{3, 3, 6, 4, 5}},
			want: 0,
		},
		{
			name: "test3",
			args: args{[]int{1, 1, 1, 3, 5}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countQuadruplets(tt.args.nums); got != tt.want {
				t.Errorf("countQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
