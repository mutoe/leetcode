package sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{4},
			want: 2,
		},
		{
			name: "test2",
			args: args{8},
			want: 2,
		},
		{
			name: "test3",
			args: args{2},
			want: 1,
		},
		{
			name: "test4",
			args: args{0},
			want: 0,
		},
		{
			name: "test5",
			args: args{1},
			want: 1,
		},
		{
			name: "test6",
			args: args{10000},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
