package count_the_number_of_consistent_strings

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}},
			want: 2,
		},
		{
			name: "test2",
			args: args{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}},
			want: 7,
		},
		{
			name: "test3",
			args: args{"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStrings(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
