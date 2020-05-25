package check_if_all_1s_are_at_least_length_k_places_away

// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away

// contest problem
// level: 2
// time: O(n) 64ms 97.14%
// space: O(1) 6.9M

// leetcode submit region begin(Prohibit modification and deletion)
func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}

	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 1 {
			break
		}
	}
	i++

	places := 0
	for ; i < len(nums); i++ {
		if nums[i] == 1 {
			if places < k {
				return false
			}
			places = 0
		} else {
			places++
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
