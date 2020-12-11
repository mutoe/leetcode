package search_insert_position

// https://leetcode.com/problems/search-insert-position

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 3.1M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}

	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// leetcode submit region end(Prohibit modification and deletion)
