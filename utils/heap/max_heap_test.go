package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		h := MaxHeap{}
		if size := h.size(); size != 0 {
			t.Errorf("should have size with 0")
		}

		h.add(1)
		h.add(5)
		h.add(3)
		h.add(2)
		h.add(10)
		h.add(6)

		if size := h.size(); size != 6 {
			t.Errorf("should have size with 6, but got %d", h.size())
		}
		if expectSort := []int{10, 5, 6, 1, 2, 3}; !reflect.DeepEqual(expectSort, h.list) {
			t.Errorf("should have correct order, expect %v but got %v", expectSort, h.list)
		}

		result := make([]int, 6)
		for i := 0; i < len(result); i++ {
			result[i] = h.extractMax()
		}
		if h.size() != 0 {
			t.Errorf("should have size with 0, but got %d", h.size())
		}
		if expectSort := []int{10, 6, 5, 3, 2, 1}; !reflect.DeepEqual(expectSort, result) {
			t.Errorf("should have correct order, expect %v but got %v", expectSort, result)
		}
	})

	t.Run("heapify", func(t *testing.T) {
		list := []int{10, 4, 22, 51, 2, 3, 3, 7, 14}
		h := MaxHeap{list}
		h.heapify()
		expectSort := []int{51, 14, 22, 10, 2, 3, 3, 7, 4}
		if !reflect.DeepEqual(expectSort, h.list) {
			t.Errorf("should have correct order, expect %v but got %v", expectSort, h.list)
		}
	})

	t.Run("replace", func(t *testing.T) {
		list := []int{51, 14, 22, 10, 2, 3, 3, 7, 4}
		h := MaxHeap{list}

		max := h.replace(5)

		if max != 51 {
			t.Errorf("return value should be 51, but got %d", max)
		}
		expectSort := []int{22, 14, 5, 10, 2, 3, 3, 7, 4}
		if !reflect.DeepEqual(expectSort, h.list) {
			t.Errorf("should have correct order, expect %v but got %v", expectSort, h.list)
		}
	})
}
