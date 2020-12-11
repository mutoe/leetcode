package find_smallest_letter_greater_than_target

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test1",
			args: args{[]byte{'c', 'f', 'j'}, 'a'},
			want: 'c',
		},
		{
			name: "test2",
			args: args{[]byte{'c', 'f', 'j'}, 'c'},
			want: 'f',
		},
		{
			name: "test3",
			args: args{[]byte{'c', 'f', 'j'}, 'd'},
			want: 'f',
		},
		{
			name: "test4",
			args: args{[]byte{'c', 'f', 'j'}, 'g'},
			want: 'j',
		},
		{
			name: "test5",
			args: args{[]byte{'c', 'f', 'j'}, 'j'},
			want: 'c',
		},
		{
			name: "test6",
			args: args{[]byte{'c', 'f', 'j'}, 'k'},
			want: 'c',
		},
		{
			name: "test7",
			args: args{[]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'n'},
			want: 'e',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
