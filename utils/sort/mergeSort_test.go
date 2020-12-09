package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	expect := []int{1, 2, 3, 4, 5}

	MergeSort(arr)

	if !reflect.DeepEqual(expect, arr) {
		t.Errorf("expect %v, but got %v", expect, arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		arr := []int{3, 2, 1, 5, 4}
		MergeSort(arr)
	}
}
