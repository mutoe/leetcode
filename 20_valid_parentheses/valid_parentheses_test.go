package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "'()' is valid",
			args: args{"()"},
			want: true,
		},
		{
			name: "'()[]{}' is valid",
			args: args{"()[]{}"},
			want: true,
		},
		{
			name: "'(]' is not valid",
			args: args{"(]"},
			want: false,
		},
		{
			name: "'([)]' is not valid",
			args: args{"([)]"},
			want: false,
		},
		{
			name: "'{[]}' is valid",
			args: args{"{[]}"},
			want: true,
		},
		{
			name: "']' is not valid",
			args: args{"]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
