package power_of_three

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{27},
			want: true,
		},
		{
			name: "test",
			args: args{0},
			want: false,
		},
		{
			name: "test",
			args: args{9},
			want: true,
		},
		{
			name: "test",
			args: args{45},
			want: false,
		},
		{
			name: "test",
			args: args{1162261468},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
