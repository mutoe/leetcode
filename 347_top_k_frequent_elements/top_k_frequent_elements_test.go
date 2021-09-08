package top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, 1, 1, 2, 2, 3}, 2},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{[]int{1}, 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
