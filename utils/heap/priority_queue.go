package heap

import "container/heap"

type Item struct {
	num int
	p   int
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
