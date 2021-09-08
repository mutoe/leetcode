package top_k_frequent_elements

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements

// level: 2
// time: O(n) 12ms 87.27%
// space: O(n) 5.4M 97.09%

//leetcode submit region begin(Prohibit modification and deletion)

type Item struct {
	num int
	p   int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].p < pq[j].p
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

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	pq := &PriorityQueue{}
	for num, count := range countMap {
		item := &Item{num, count}
		if pq.Len() < k {
			heap.Push(pq, item)
		} else if count > pq.Peek().p {
			heap.Pop(pq)
			heap.Push(pq, item)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = (*pq)[i].num
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
