package find_greatest_common_divisor_of_array

// https://leetcode.com/problems/find-greatest-common-divisor-of-array

// level: 1
// time: O(n) 4ms 90.98%
// space: 3.1M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
	}
	for min > 0 {
		min, max = max%min, min
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
