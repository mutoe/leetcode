package find_k_pairs_with_smallest_sums

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums

// level: 2

//leetcode submit region begin(Prohibit modification and deletion)
import "container/heap"

type Item struct {
	num1 int
	num2 int
	p    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].p >= pq[j].p
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &PriorityQueue{}
	for _, i := range nums1 {
		for _, j := range nums2 {
			if h.Len() < k {
				heap.Push(h, &Item{i, j, i + j})
			} else if sum := i + j; sum < h.Peek().p {
				heap.Pop(h)
				heap.Push(h, &Item{i, j, sum})
			}
		}
	}
	if h.Len() < k {
		k = h.Len()
	}
	ret := make([][]int, k)
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(h).(*Item)
		ret[i] = []int{item.num1, item.num2}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
