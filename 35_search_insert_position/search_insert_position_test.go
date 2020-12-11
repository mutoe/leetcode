package search_insert_position

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 3, 5, 6}, 5},
			want: 2,
		},
		{
			name: "test2",
			args: args{[]int{1, 3, 5, 6}, 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{[]int{1, 3, 5, 6}, 7},
			want: 4,
		},
		{
			name: "test4",
			args: args{[]int{1, 3, 5, 6}, 0},
			want: 0,
		},
		{
			name: "test5",
			args: args{[]int{1}, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
