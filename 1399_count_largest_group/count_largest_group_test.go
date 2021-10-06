package count_largest_group

import "testing"

func Test_countLargestGroup(t *testing.T) {
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
			args: args{13},
			want: 4,
		},
		{
			name: "test2",
			args: args{2},
			want: 2,
		},
		{
			name: "test3",
			args: args{15},
			want: 6,
		},
		{
			name: "test4",
			args: args{24},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLargestGroup(tt.args.n); got != tt.want {
				t.Errorf("countLargestGroup(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
