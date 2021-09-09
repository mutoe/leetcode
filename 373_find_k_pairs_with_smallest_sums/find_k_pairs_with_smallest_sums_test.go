package find_k_pairs_with_smallest_sums

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				[]int{1, 7, 11},
				[]int{2, 4, 6},
				3,
			},
			want: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			name: "test2",
			args: args{
				[]int{1, 1, 2},
				[]int{1, 2, 3},
				2,
			},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "test3",
			args: args{
				[]int{1, 2},
				[]int{3},
				3,
			},
			want: [][]int{{1, 3}, {2, 3}},
		},
		{
			name: "test4",
			args: args{
				[]int{-13, 23, 44, 117, 900, 990},
				[]int{-15, 20, 35, 118, 223, 500, 663, 717, 789, 813},
				10,
			},
			want: [][]int{{-13, -15}, {-13, 20}, {23, -15}, {-13, 35}, {44, -15}, {23, 20}, {23, 35}, {44, 20}, {44, 35}, {117, -15}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
