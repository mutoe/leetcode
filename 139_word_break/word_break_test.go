package word_break

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "test1",
		//	args: args{"leetcode", []string{"leet", "code"}},
		//	want: true,
		//},
		//{
		//	name: "test2",
		//	args: args{"applepenapple", []string{"apple", "pen"}},
		//	want: true,
		//},
		//{
		//	name: "test3",
		//	args: args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
		//	want: false,
		//},
		//{
		//	name: "test4",
		//	args: args{"aaaaaaa", []string{"aaaa", "aaa"}},
		//	want: true,
		//},
		//{
		//	name: "test5",
		//	args: args{"bb", []string{"a", "b", "bbb", "bbbb"}},
		//	want: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
