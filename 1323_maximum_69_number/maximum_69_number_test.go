package maximum_69_number

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{9669},
			want: 9969,
		},
		{
			name: "test2",
			args: args{9996},
			want: 9999,
		},
		{
			name: "test3",
			args: args{9999},
			want: 9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
