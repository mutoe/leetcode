package increasing_decreasing_string

import "testing"

func Test_sortString(t *testing.T) {
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
			args: args{s: "aaaabbbbcccc"},
			want: "abccbaabccba",
		},
		{
			name: "test2",
			args: args{s: "rat"},
			want: "art",
		},
		{
			name: "test3",
			args: args{s: "leetcode"},
			want: "cdelotee",
		},
		{
			name: "test4",
			args: args{s: "ggggggg"},
			want: "ggggggg",
		},
		{
			name: "test5",
			args: args{s: "spo"},
			want: "ops",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
