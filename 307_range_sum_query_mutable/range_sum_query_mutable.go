package range_sum_query_mutable

import "math"

// https://leetcode.com/problems/range-sum-query-mutable

// level: 2
// time: 592ms 76.47%
// space: 22.6M 39.71%

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
	st.tree[index] = st.merge(st.tree[leftIndex], st.tree[rightIndex])
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
	return st.merge(left, right)
}

func (st *NumArray) Query(l, r int) int {
	if l < 0 || l >= len(st.data) || r < 0 || r >= len(st.data) || l > r {
		return math.MinInt64
	}
	return st.query(0, 0, len(st.data)-1, l, r)
}

func (st *NumArray) merge(l, r int) int {
	return l + r
}

func (st *NumArray) update(rootIndex, l, r, index, val int) {
	if l == r {
		st.tree[rootIndex] = val
		return
	}
	mid := l + (r-l)/2
	leftIndex, rightIndex := rootIndex*2+1, rootIndex*2+2
	if index >= mid+1 {
		st.update(rightIndex, mid+1, r, index, val)
	} else {
		st.update(leftIndex, l, mid, index, val)
	}

	st.tree[rootIndex] = st.merge(st.tree[leftIndex], st.tree[rightIndex])
}

func (st *NumArray) Update(index int, val int) {
	st.data[index] = val
	st.update(0, 0, len(st.data)-1, index, val)
}

func (st *NumArray) SumRange(left int, right int) int {
	return st.Query(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
