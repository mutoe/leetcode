package word_search_ii

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, []string{"oath", "pea", "eat", "rain"}},
			want: []string{"eat", "oath"},
		},
		{
			name: "test2",
			args: args{[][]byte{
				{'a', 'b'},
				{'c', 'd'},
			}, []string{"abcd"}},
			want: []string{},
		},
		{
			name: "test3",
			args: args{[][]byte{{'a'}}, []string{"a"}},
			want: []string{"a"},
		},
		{
			name: "test4",
			args: args{[][]byte{
				{'o', 'a', 'b', 'n'},
				{'o', 't', 'a', 'e'},
				{'a', 'h', 'k', 'r'},
				{'a', 'f', 'l', 'v'},
			}, []string{"oa", "oaa"}},
			want: []string{"oa", "oaa"},
		},
		{
			name: "test5",
			args: args{[][]byte{{'a', 'a'}}, []string{"aaa"}},
			want: []string{},
		},
		{
			name: "test6",
			args: args{[][]byte{
				{'a', 'b'},
				{'c', 'd'},
			}, []string{"abdca"}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.args.board, tt.args.words)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
