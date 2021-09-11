package find_k_closest_elements

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/find-k-closest-elements

// level: 2
// time: O(n*log(n)) 123ms 6.82%
// space: O(n) 7.5M 18.94%

//leetcode submit region begin(Prohibit modification and deletion)
type Item struct {
	number int
	p      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].p > pq[j].p
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old) - 1
	x := old[n]
	*pq = old[0:n]
	return x
}

func (pq PriorityQueue) Peek() Item {
	return *pq[0]
}

func (pq *PriorityQueue) Init(nums []*Item) {
	*pq = nums[:]
	heap.Init(pq)
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	if n == k {
		return arr
	}
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[n-1] {
		return arr[n-k : n]
	}
	h := &PriorityQueue{}
	for i := 0; i < len(arr); i++ {
		delta := arr[i] - x
		if delta < 0 {
			delta *= -1
		}
		if h.Len() < k {
			heap.Push(h, &Item{arr[i], delta})
		} else if delta < h.Peek().p {
			heap.Pop(h)
			heap.Push(h, &Item{arr[i], delta})
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = (*h)[i].number
	}
	sort.Ints(ret)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
