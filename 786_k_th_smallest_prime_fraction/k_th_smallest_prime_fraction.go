package k_th_smallest_prime_fraction

// https://leetcode.com/problems/k-th-smallest-prime-fraction

// level: 3

// max heap solution
// time: O(n^2) 486ms 50%
// space: O(n) 31.8M 100%

//leetcode submit region begin(Prohibit modification and deletion)

import "container/heap"

type Item struct {
	numerator   int
	denominator int
	fraction    float32
}

type FloatHeap []*Item

func (h FloatHeap) Len() int {
	return len(h)
}

func (h FloatHeap) Less(i, j int) bool {
	return h[i].fraction > h[j].fraction
}

func (h FloatHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func (h FloatHeap) Peek() Item {
	return *h[0]
}

func (h *FloatHeap) Init(nums []*Item) {
	*h = nums[:]
	heap.Init(h)
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	if len(arr) == 2 {
		return arr
	}
	h := &FloatHeap{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fraction := float32(arr[i]) / float32(arr[j])
			item := &Item{arr[i], arr[j], fraction}
			if h.Len() < k {
				heap.Push(h, item)
			} else if fraction < h.Peek().fraction {
				heap.Pop(h)
				heap.Push(h, item)
			}
		}
	}
	return []int{h.Peek().numerator, h.Peek().denominator}
}

//leetcode submit region end(Prohibit modification and deletion)
