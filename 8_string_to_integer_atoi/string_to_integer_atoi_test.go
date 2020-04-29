package string_to_integer_atoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{"42"},
			42,
		},
		{
			args: args{"     -42"},
			want: -42,
		},
		{
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			args: args{"words and 987"},
			want: 0,
		},
		{
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			args: args{"52s3"},
			want: 52,
		},
		{
			args: args{"+52s3"},
			want: 52,
		},
		{
			args: args{"++52s3"},
			want: 0,
		},
		{
			args: args{"-+52s3"},
			want: 0,
		},
		{
			args: args{"+0 123"},
			want: 0,
		},
		{
			args: args{"9223372036854775808"},
			want: 2147483647,
		},
		{
			args: args{"2147483800"},
			want: 2147483647,
		},
		{
			args: args{"0-1"},
			want: 0,
		},
		{
			args: args{"-01"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.str, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi(%s) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
