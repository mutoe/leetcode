package perfect_number

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
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
			want: false,
		},
		{
			name: "test",
			args: args{2},
			want: false,
		},
		{
			name: "test",
			args: args{3},
			want: false,
		},
		{
			name: "test",
			args: args{4},
			want: false,
		},
		{
			name: "test",
			args: args{5},
			want: false,
		},
		{
			name: "test",
			args: args{6},
			want: true,
		},
		{
			name: "test",
			args: args{28},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
