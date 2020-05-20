package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "test1.1",
			args: args{"abcad"},
			want: 4,
		},
		{
			name: "test2",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "test3",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "test4",
			args: args{""},
			want: 0,
		},
		{
			name: "test5",
			args: args{"abcb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
