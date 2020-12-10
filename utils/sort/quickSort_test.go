package sort

import (
	"fmt"
	"reflect"
	"testing"

	".."
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{3, 2, 1, 5, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			args: []int{0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			args: []int{8, 1, 2, 1, 5},
			want: []int{1, 1, 2, 5, 8},
		},
		{
			args: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			args: []int{5, 4, 4, 3, 1, 2, 5, 2, 6, 2},
			want: []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 6},
		},
		{
			args: []int{1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args), func(t *testing.T) {
			QuickSort(tt.args)
			if !reflect.DeepEqual(tt.want, tt.args) {
				t.Errorf("expect %v, but got %v", tt.want, tt.args)
			}
		})
	}
}

func BenchmarkQuickSortOrdered(b *testing.B) {
	const n = 1e5
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateOrderedArray(n)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSortRandom(b *testing.B) {
	const n = 1e5
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateRandomArray(n, n)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSortEquivalent(b *testing.B) {
	const n = 1e5
	arr := utils.GenerateRandomArray(n, 1)
	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}
