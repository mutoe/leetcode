package range_sum_query_immutable

import "math"

// https://leetcode.com/problems/range-sum-query-immutable

// level: 1
// time: O(n*log(n)) 52ms 36.51%
// space: O(n) 8.6M 94.78%

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	data []int
	tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)*4)
	numArray := NumArray{nums, tree}
	numArray.build(0, 0, len(nums)-1)
	return numArray
}

func (st *NumArray) build(index, l, r int) {
	if l == r {
		st.tree[index] = st.data[l]
		return
	}
	mid := l + (r-l)/2
	leftIndex, rightIndex := index*2+1, index*2+2
	st.build(leftIndex, l, mid)
	st.build(rightIndex, mid+1, r)
	st.tree[index] = st.Merge(st.tree[leftIndex], st.tree[rightIndex])
}

func (st *NumArray) query(index, l, r, qL, qR int) int {
	if l == qL && r == qR {
		return st.tree[index]
	}
	mid := l + (r-l)/2
	leftIndex, rightIndex := index*2+1, index*2+2
	if qL >= mid+1 {
		return st.query(rightIndex, mid+1, r, qL, qR)
	}
	if qR <= mid {
		return st.query(leftIndex, l, mid, qL, qR)
	}
	left := st.query(leftIndex, l, mid, qL, mid)
	right := st.query(rightIndex, mid+1, r, mid+1, qR)
	return st.Merge(left, right)
}

func (st *NumArray) Query(l, r int) int {
	if l < 0 || l >= len(st.data) || r < 0 || r >= len(st.data) || l > r {
		return math.MinInt64
	}
	return st.query(0, 0, len(st.data)-1, l, r)
}

func (st *NumArray) Merge(l, r int) int {
	return l + r
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
