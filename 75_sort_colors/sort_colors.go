package sort_colors

// https://leetcode.com/problems/sort-colors

// level: 2
// time: O(n) 0ms 100%
// space: O(1) 2.0M 64.29%

// One-pass solution (three-pointer)

// leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	zero, i, two := -1, 0, len(nums)
	for i < two {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		} else {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		}
	}

}

// leetcode submit region end(Prohibit modification and deletion)
