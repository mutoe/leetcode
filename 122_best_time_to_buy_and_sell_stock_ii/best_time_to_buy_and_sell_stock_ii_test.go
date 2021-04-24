package best_time_to_buy_and_sell_stock_ii

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "test3",
			args: args{[]int{7, 6, 4, 3, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
