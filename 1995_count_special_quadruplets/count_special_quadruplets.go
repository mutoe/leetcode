package count_special_quadruplets

// https://leetcode.com/problems/count-special-quadruplets

// level: 1
// time: O(n^4) 12ms 82.41%
// space: O(1) 2.3M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func countQuadruplets(nums []int) int {
	n := len(nums)
	ret := 0
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						ret++
					}
				}
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
