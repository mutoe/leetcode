package count_primes

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{0},
			want: 0,
		},
		{
			name: "test",
			args: args{1},
			want: 0,
		},
		{
			name: "test",
			args: args{2},
			want: 0,
		},
		{
			name: "test",
			args: args{3},
			want: 1,
		},
		{
			name: "test",
			args: args{10},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
