package reverse_vowels_of_a_string

import "testing"

func Test_reverseVowels(t *testing.T) {
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
			args: args{"hello"},
			want: "holle",
		},
		{
			name: "test2",
			args: args{"leetcode"},
			want: "leotcede",
		},
		{
			name: "test3",
			args: args{"HOelloaab123"},
			want: "HaalloeOb123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
