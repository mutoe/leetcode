package sort

import (
	"reflect"
	"testing"

	".."
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	expect := []int{1, 2, 3, 4, 5}

	MergeSort(arr)

	if !reflect.DeepEqual(expect, arr) {
		t.Errorf("expect %v, but got %v", expect, arr)
	}
}

func BenchmarkMergeSortOrdered(b *testing.B) {
	const n = 1e5
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateOrderedArray(n)
		b.StartTimer()
		MergeSort(arr)
	}
}

func BenchmarkMergeSortRandom(b *testing.B) {
	const n = 1e5
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateRandomArray(n, n)
		b.StartTimer()
		MergeSort(arr)
	}
}

func BenchmarkMergeSortEquivalent(b *testing.B) {
	const n = 1e5
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := utils.GenerateRandomArray(n, 1)
		b.StartTimer()
		MergeSort(arr)
	}
}
