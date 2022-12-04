package contains_duplicate_ii

// https://leetcode.com/problems/contains-duplicate-ii

// level: 1
// time: O(n) 250ms 37.93%
// space: O(n) 13.7M 29.45%

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	// key: num, value: previous index
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if m[num] == 0 {
			m[num] = i + 1
		} else {
			if i+1-m[num] <= k {
				return true
			}
			m[num] = i + 1
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
