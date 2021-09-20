package range_sum_query_mutable

import "math"

// https://leetcode.com/problems/range-sum-query-mutable

// level: 2
// time: 584ms 82.81%
// space: 22.1M 51.56%

// Sqrt Decomposition solution

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	data   []int
	N      int
	blocks []int
	B      int
	Bn     int
}

func Constructor(nums []int) NumArray {
	N := len(nums)
	sqrt := math.Sqrt(float64(N))
	blockSize := int(sqrt)
	blockCount := int(math.Ceil(float64(N) / float64(blockSize)))

	block := make([]int, blockCount)
	for i := 0; i < N; i++ {
		block[i/blockSize] += nums[i]
	}

	return NumArray{
		data:   nums,
		N:      N,
		blocks: block,
		B:      blockSize,
		Bn:     blockCount,
	}
}

func (sqrt *NumArray) Update(index int, val int) {
	oldValue := sqrt.data[index]
	sqrt.data[index] = val
	sqrt.blocks[index/sqrt.B] += val - oldValue
}

func (sqrt *NumArray) SumRange(a, b int) int {
	bs, be := a/sqrt.B, b/sqrt.B
	res := int(0)
	if bs == be {
		for i := a; i <= b; i++ {
			res += sqrt.data[i]
		}
		return res
	}
	for i := a; i < sqrt.B*(bs+1); i++ {
		res += sqrt.data[i]
	}
	for b := bs + 1; b < be; b++ {
		res += sqrt.blocks[b]
	}
	for i := be * sqrt.B; i <= b; i++ {
		res += sqrt.data[i]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
