package maximum_product_of_two_elements_in_an_array

import "container/heap"

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array

// level: 1
// time: O(n) 4ms 91.3%
// space: O(n) 3M 15.22%

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
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func (h *IntHeap) Init(nums []int) {
	*h = nums[:]
	heap.Init(h)
}

func maxProduct(nums []int) int {
	h := &IntHeap{}
	h.Init(nums)
	a := heap.Pop(h).(int)
	b := (*h)[0]
	return (a - 1) * (b - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
