package make_the_string_great

import "testing"

func Test_makeGood(t *testing.T) {
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
			args: args{"leEeetcode"},
			want: "leetcode",
		},
		{
			name: "test2",
			args: args{"abBAcC"},
			want: "",
		},
		{
			name: "test3",
			args: args{"s"},
			want: "s",
		},
		{
			name: "test4",
			args: args{"mC"},
			want: "mC",
		},
		{
			name: "test5",
			args: args{"Pp"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
