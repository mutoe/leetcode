package find_the_middle_index_in_array

import "testing"

func Test_findMiddleIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{2, 3, -1, 8, 4}},
			want: 3,
		},
		{
			name: "test2",
			args: args{[]int{1, -1, 4}},
			want: 2,
		},
		{
			name: "test3",
			args: args{[]int{2, 5}},
			want: -1,
		},
		{
			name: "test4",
			args: args{[]int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
