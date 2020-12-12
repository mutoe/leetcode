package capacity_to_ship_packages_within_d_days

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		D       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
			want: 15,
		},
		{
			name: "test2",
			args: args{[]int{3, 2, 2, 4, 1, 4}, 3},
			want: 6,
		},
		{
			name: "test3",
			args: args{[]int{1, 2, 3, 1, 1}, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.D); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDays(t *testing.T) {
	type args struct {
		weights []int
		weight  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 1, 1}, 4},
			want: 3,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 3, 1, 1}, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDays(tt.args.weights, tt.args.weight); got != tt.want {
				t.Errorf("getDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
