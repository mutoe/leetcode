package longest_happy_prefix

import "testing"

func Test_longestPrefix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"level"},
			want: "l",
		},
		{
			name: "test2",
			args: args{"ababab"},
			want: "abab",
		},
		{
			name: "test3",
			args: args{"leetcodeleet"},
			want: "leet",
		},
		{
			name: "test4",
			args: args{"a"},
			want: "",
		},
		{
			name: "test5",
			args: args{"aaaaaaa"},
			want: "aaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrefix(tt.args.s); got != tt.want {
				t.Errorf("longestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
