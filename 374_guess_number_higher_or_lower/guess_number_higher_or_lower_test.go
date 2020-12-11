package guess_number_higher_or_lower

import "testing"

func Test_guessNumber(t *testing.T) {
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
			args: args{10},
			want: 6,
		},
		{
			name: "test2",
			args: args{1},
			want: 1,
		},
		{
			name: "test3",
			args: args{2},
			want: 1,
		},
		{
			name: "test4",
			args: args{2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.want
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
