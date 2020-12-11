package first_bad_version

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{5},
			want: 4,
		},
		{
			name: "test2",
			args: args{1},
			want: 1,
		},
		{
			name: "test3",
			args: args{3},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVersion = tt.want
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
