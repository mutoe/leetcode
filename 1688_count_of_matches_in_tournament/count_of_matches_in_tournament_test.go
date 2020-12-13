package count_of_matches_in_tournament

import "testing"

func Test_numberOfMatches(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{7},
			want: 6,
		},
		{
			args: args{14},
			name: "test2",
			want: 13,
		},
		{
			name: "test3",
			args: args{3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatches(tt.args.n); got != tt.want {
				t.Errorf("numberOfMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
