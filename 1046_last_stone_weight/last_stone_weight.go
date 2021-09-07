package last_stone_weight

import "container/heap"

// https://leetcode.com/problems/last-stone-weight

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2M 46.15%

//leetcode submit region begin(Prohibit modification and deletion)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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
	x := old[last]
	*h = old[0:last]
	return x
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func (h *IntHeap) Init(nums []int) {
	*h = nums[:]
	heap.Init(h)
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	h.Init(stones)

	for {
		if h.Len() == 0 {
			return 0
		} else if h.Len() == 1 {
			return h.Peek()
		} else {
			a, b := heap.Pop(h).(int), heap.Pop(h).(int)
			delta := a - b
			if delta < 0 {
				delta = -delta
			}
			if delta != 0 {
				heap.Push(h, delta)
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
