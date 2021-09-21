package longest_chunked_palindrome_decomposition

import "testing"

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{"ghiabcdefhelloadamhelloabcdefghi"},
			want: 7,
		},
		{
			name: "test2",
			args: args{"merchant"},
			want: 1,
		},
		{
			name: "test3",
			args: args{"antaprezatepzapreanta"},
			want: 11,
		},
		{
			name: "test4",
			args: args{"aaa"},
			want: 3,
		},
		{
			name: "test5",
			args: args{"elvtoelvto"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}
