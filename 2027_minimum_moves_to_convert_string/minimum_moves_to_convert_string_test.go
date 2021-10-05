package minimum_moves_to_convert_string

import "testing"

func Test_minimumMoves(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{"XXX"},
			want: 1,
		},
		{
			name: "test2",
			args: args{"XXOX"},
			want: 2,
		},
		{
			name: "test3",
			args: args{"OOOO"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.s); got != tt.want {
				t.Errorf("minimumMoves(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
