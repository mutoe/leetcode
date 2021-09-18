package range_sum_query_mutable

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		mode     int
		boundary []int
		want     int
	}
	tests := []struct {
		name      string
		nums      []int
		testcases []args
	}{
		//{
		//	name: "test1",
		//	nums: []int{1, 3, 5},
		//	testcases: []args{
		//		{1, []int{0, 2}, 9},
		//		{2, []int{1, 2}, 0},
		//		{1, []int{0, 2}, 8},
		//	},
		//},
		{
			name: "test2",
			nums: []int{0, 9, 5, 7, 3},
			testcases: []args{
				{1, []int{4, 4}, 3},
				{1, []int{2, 4}, 15},
				{1, []int{3, 3}, 7},
				{2, []int{4, 5}, 0},
				{2, []int{1, 7}, 0},
				{2, []int{0, 8}, 0},
				{1, []int{1, 2}, 12},
				{2, []int{1, 9}, 0},
				{1, []int{4, 4}, 5},
				{2, []int{3, 4}, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(tt.nums)
			for _, testcase := range tt.testcases {
				if testcase.mode == 1 {
					left, right, want := testcase.boundary[0], testcase.boundary[1], testcase.want
					if got := obj.SumRange(left, right); got != want {
						t.Errorf("SumRange(%v, %v) = %v, want %v", left, right, got, want)
					}
				} else {
					index, value := testcase.boundary[0], testcase.boundary[1]
					obj.Update(index, value)
				}
			}
		})
	}
}
