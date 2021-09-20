package first_unique_character_in_a_string

import "testing"

func Test_firstUniqChar(t *testing.T) {
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
			args: args{"leetcode"},
			want: 0,
		},
		{
			name: "test2",
			args: args{"loveleetcode"},
			want: 2,
		},
		{
			name: "test3",
			args: args{"aaabb"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
