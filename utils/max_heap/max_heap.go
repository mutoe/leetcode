package max_heap

type MaxHeap struct {
	list []int
}

func (h MaxHeap) size() int {
	return len(h.list)
}

func (h MaxHeap) parent(i int) int {
	if i <= 0 {
		return 0
	}
	return (i - 1) / 2
}

func (h MaxHeap) leftChild(i int) int {
	return i*2 + 1
}

func (h MaxHeap) rightChild(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) add(n int) {
	h.list = append(h.list, n)
	cur := h.size() - 1
	for par := h.parent(cur); par >= 0 && h.list[cur] > h.list[par]; par = h.parent(cur) {
		h.list[cur], h.list[par] = h.list[par], h.list[cur]
		cur = par
	}
}

func (h *MaxHeap) siftDown(cur int) {
	for left := h.leftChild(cur); h.leftChild(cur) < h.size(); left = h.leftChild(cur) {
		right := left + 1
		max := left

		if right < h.size() && h.list[right] > h.list[left] {
			max = right
		}
		if h.list[cur] >= h.list[max] {
			break
		}
		h.list[cur], h.list[max] = h.list[max], h.list[cur]
		cur = max
	}
}

func (h *MaxHeap) extractMax() (result int) {
	result = h.list[0]
	h.list = append([]int{h.list[h.size()-1]}, h.list[1:h.size()-1]...)
	h.siftDown(0)
	return
}

func (h *MaxHeap) replace(newVal int) (result int) {
	result = h.list[0]
	h.list[0] = newVal
	h.siftDown(0)
	return
}

func (h *MaxHeap) heapify() {
	list := h.list[:]
	for cur := len(list) / 2; cur >= 0; cur-- {
		h.siftDown(cur)
	}
}
