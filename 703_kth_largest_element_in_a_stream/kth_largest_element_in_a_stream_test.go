package kth_largest_element_in_a_stream

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name string
		k    int
		init []int
		args []int
		want []int
	}{
		{
			name: "test1",
			k:    3,
			init: []int{4, 5, 8, 2},
			args: []int{3, 5, 10, 9, 4},
			want: []int{4, 5, 5, 8, 8},
		},
		{
			name: "test2",
			k:    1,
			init: []int{},
			args: []int{-3, -2, -4, 0, 4},
			want: []int{-3, -2, -2, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kth := Constructor(tt.k, tt.init)
			for i := 0; i < len(tt.args); i++ {
				if got := kth.Add(tt.args[i]); got != tt.want[i] {
					t.Errorf("%d: Add(%d) = %v, want %v", i, tt.args[i], got, tt.want[i])
				}
			}
		})
	}
}
