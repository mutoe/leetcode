package kth_largest_element_in_a_stream

import "container/heap"

// https://leetcode.com/problems/kth-largest-element-in-a-stream

// level: 1
// time: O(n*log(n)) 32ms 78.22%
// space: O(n) 8.2M 63.37%

//leetcode submit region begin(Prohibit modification and deletion)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	last := len(old) - 1
	ret := old[last]
	*h = old[0:last]
	return ret
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		&IntHeap{},
		k,
	}

	for _, num := range nums {
		kth.Add(num)
	}

	return kth
}

func (kth *KthLargest) Add(val int) int {
	if kth.h.Len() < kth.k {
		heap.Push(kth.h, val)
	} else if val > kth.h.Peek() {
		heap.Pop(kth.h)
		heap.Push(kth.h, val)
	}
	return kth.h.Peek()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
