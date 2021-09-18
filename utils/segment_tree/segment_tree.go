package segment_tree

import "math"

type NodeType int

type SegmentTree struct {
	data []NodeType
	tree []NodeType
}

func InitSegmentTree(nums []NodeType) SegmentTree {
	tree := make([]NodeType, len(nums)*4)
	numArray := SegmentTree{nums, tree}
	numArray.build(0, 0, len(nums)-1)
	return numArray
}

func (st *SegmentTree) build(index, l, r int) {
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

func (st *SegmentTree) query(index, l, r, qL, qR int) NodeType {
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

func (st *SegmentTree) Query(l, r int) NodeType {
	if l < 0 || l >= len(st.data) || r < 0 || r >= len(st.data) || l > r {
		return math.MinInt64
	}
	return st.query(0, 0, len(st.data)-1, l, r)
}

func (st *SegmentTree) merge(l, r NodeType) NodeType {
	return l + r
}

//func (st *SegmentTree) SumRange(left int, right int) int {
//	return st.query(left, right)
//}
