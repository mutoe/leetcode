package ugly_number

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{0},
			want: false,
		},
		{
			name: "test",
			args: args{1},
			want: true,
		},
		{
			name: "test",
			args: args{2},
			want: true,
		},
		{
			name: "test",
			args: args{5},
			want: true,
		},
		{
			name: "test",
			args: args{6},
			want: true,
		},
		{
			name: "test",
			args: args{8},
			want: true,
		},
		{
			name: "test",
			args: args{14},
			want: false,
		},
		{
			name: "test",
			args: args{-8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
