package container_with_most_water

// https://leetcode.com/problems/container-with-most-water

// level: 2
// time: O(n^2) 296ms 26.72%
// space: O(1) 5.8M

// leetcode submit region begin(Prohibit modification and deletion)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) (ret int) {
	size := len(height)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if area := (j - i) * min(height[i], height[j]); area > ret {
				ret = area
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
