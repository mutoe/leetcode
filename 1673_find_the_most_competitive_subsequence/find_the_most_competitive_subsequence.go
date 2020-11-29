package find_the_most_competitive_subsequence

import "sort"

// https://leetcode.com/problems/find-the-most-competitive-subsequence

// level: 1
// time: O(n^2) 180ms 0%
// space: O(n) 10.1M 0%

// leetcode submit region begin(Prohibit modification and deletion)
func mostCompetitive(nums []int, k int) []int {
	var ret []int
	l, r := 0, len(nums)-k

	if sort.IntsAreSorted(nums) {
		return nums[0:k]
	}

	for len(ret) < k && r < len(nums) {
		minIndex := min(nums[l : r+1])
		ret = append(ret, nums[l+minIndex])
		l += minIndex + 1
		r++
	}
	return ret
}

func min(nums []int) int {
	min := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[min] {
			min = i
		}
	}
	return min
}

// leetcode submit region end(Prohibit modification and deletion)
