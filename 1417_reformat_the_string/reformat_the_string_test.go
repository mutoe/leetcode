package reformat_the_string

import "testing"

func Test_reformat(t *testing.T) {
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
			args: args{"a0b1c2"},
			want: "0a1b2c",
		},
		{
			name: "test2",
			args: args{"leetcode"},
			want: "",
		},
		{
			name: "test3",
			args: args{"1229857369"},
			want: "",
		},
		{
			name: "test4",
			args: args{"covid2019"},
			want: "c2o0v1i9d",
		},
		{
			name: "test5",
			args: args{"ab123"},
			want: "1a2b3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.args.s); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
