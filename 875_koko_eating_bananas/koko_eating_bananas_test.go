package koko_eating_bananas

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{3, 6, 7, 11}, 8},
			want: 4,
		},
		{
			name: "test2",
			args: args{[]int{30, 11, 23, 4, 20}, 5},
			want: 30,
		},
		{
			name: "test3",
			args: args{[]int{30, 11, 23, 4, 20}, 6},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				// t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
