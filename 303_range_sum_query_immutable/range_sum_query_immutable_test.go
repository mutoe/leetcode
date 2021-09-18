package range_sum_query_immutable

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		boundary []int
		want     int
	}
	tests := []struct {
		name      string
		nums      []int
		testcases []args
	}{
		{
			name: "test1",
			nums: []int{-2, 0, 3, -5, 2, -1},
			testcases: []args{
				{[]int{0, 2}, 1},
				{[]int{2, 5}, -1},
				{[]int{0, 5}, -3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(tt.nums)
			for _, testcase := range tt.testcases {
				left, right, want := testcase.boundary[0], testcase.boundary[1], testcase.want
				if got := obj.SumRange(left, right); got != want {
					t.Errorf("SumRange(%v, %v) = %v, want %v", left, right, got, want)
				}
			}
		})
	}
}
