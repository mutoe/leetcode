package range_sum_query_immutable

// https://leetcode.com/problems/range-sum-query-immutable

// level: 1
// time: O(n) 24ms 98.41%
// space: O(n) 9.6M 66.21%

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray{sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
