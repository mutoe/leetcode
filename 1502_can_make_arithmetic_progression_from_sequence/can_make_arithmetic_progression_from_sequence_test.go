package can_make_arithmetic_progression_from_sequence

import "testing"

func Test_canMakeArithmeticProgression(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{3, 5, 1}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 4}},
			want: false,
		},
		{
			name: "test3",
			args: args{[]int{0, 0, 0, 0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.args.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
