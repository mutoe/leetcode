package sort

import (
	"reflect"
	"testing"

	".."
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	expect := []int{1, 2, 3, 4, 5}

	QuickSort(arr)

	if !reflect.DeepEqual(expect, arr) {
		t.Errorf("expect %v, but got %v", expect, arr)
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
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateRandomArray(n, 1)
		b.StartTimer()
		QuickSort(arr)
	}
}
