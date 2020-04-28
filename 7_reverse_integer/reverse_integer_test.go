package reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
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
			args: args{123},
			want: 321,
		},
		{
			name: "test2",
			args: args{-123},
			want: -321,
		},
		{
			name: "test3",
			args: args{120},
			want: 21,
		},
		{
			name: "test4",
			args: args{1534236469},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
