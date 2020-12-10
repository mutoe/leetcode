package binary_search

// https://leetcode.com/problems/binary-search

// level: 1
// time: O(log(n)) 36ms 67.19%
// space: O(log(n)) 6.5M 81.82%

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	return binary(nums, 0, len(nums)-1, target)
}

func binary(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)>>1
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binary(nums, l, mid-1, target)
	} else {
		return binary(nums, mid+1, r, target)
	}
}

// leetcode submit region end(Prohibit modification and deletion)
