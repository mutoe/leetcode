package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	expect := []int{1, 2, 3, 4, 5}

	QuickSort(arr)

	if !reflect.DeepEqual(expect, arr) {
		t.Errorf("expect %v, but got %v", expect, arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		arr := []int{3, 2, 1, 5, 4}
		QuickSort(arr)
	}
}
