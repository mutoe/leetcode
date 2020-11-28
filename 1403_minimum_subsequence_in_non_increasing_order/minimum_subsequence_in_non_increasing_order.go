package minimum_subsequence_in_non_increasing_order

import "sort"

// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order

// level: 1
// time: O(n*log(n)) 4ms 97.44%
// space: O(n) 3.2M 100%

// leetcode submit region begin(Prohibit modification and deletion)

func Sum(q []int) int {
	sum := 0
	for _, n := range q {
		sum += n
	}
	return sum
}

func minSubsequence(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := Sum(nums)
	i := 0
	for ; i < len(nums); i++ {
		leftSum := Sum(nums[:i+1])
		if leftSum > sum-leftSum {
			break
		}
	}
	return nums[:i+1]
}

// leetcode submit region end(Prohibit modification and deletion)
