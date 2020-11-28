package shuffle_string

import "testing"

func Test_restoreString(t *testing.T) {
	type args struct {
		s       string
		indices []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s:       "codeleet",
				indices: []int{4, 5, 6, 7, 0, 2, 1, 3},
			},
			want: "leetcode",
		},
		{
			name: "test2",
			args: args{
				s:       "abc",
				indices: []int{0, 1, 2},
			},
			want: "abc",
		},
		{
			name: "test3",
			args: args{
				s:       "aiohn",
				indices: []int{3, 1, 4, 2, 0},
			},
			want: "nihao",
		},
		{
			name: "test4",
			args: args{
				s:       "aaiougrt",
				indices: []int{4, 0, 2, 6, 7, 3, 1, 5},
			},
			want: "arigatou",
		},
		{
			name: "test5",
			args: args{
				s:       "art",
				indices: []int{1, 0, 2},
			},
			want: "rat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString(tt.args.s, tt.args.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
