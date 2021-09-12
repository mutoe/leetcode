package reverse_prefix_of_word

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"abcdefd", 'd'},
			want: "dcbaefd",
		},
		{
			name: "test2",
			args: args{"xyxzxe", 'z'},
			want: "zxyxxe",
		},
		{
			name: "test3",
			args: args{"abcd", 'z'},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
