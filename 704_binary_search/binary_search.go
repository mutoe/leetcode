package binary_search

// https://leetcode.com/problems/binary-search

// level: 1
// time: O(log(n)) 36ms 67.19%
// space: O(1) 6.5M 81.82%

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
