package sort_colors

// https://leetcode.com/problems/sort-colors

// level: 2
// time: O(n) 0ms 100%
// space: O(1) 2.1M 64.29%

// leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	colors := [3]int{}
	for _, num := range nums {
		colors[num]++
	}
	k := 0
	for n, count := range colors {
		for i := k; i < k+count; i++ {
			nums[i] = n
		}
		k += count
	}
}

// leetcode submit region end(Prohibit modification and deletion)
