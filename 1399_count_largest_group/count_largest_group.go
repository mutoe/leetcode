package count_largest_group

// https://leetcode.com/problems/count-largest-group

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.1M 92.31%

//leetcode submit region begin(Prohibit modification and deletion)
func countLargestGroup(n int) int {
	if n < 10 {
		return n
	}

	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		cur := i
		sum := 0
		for cur > 0 {
			sum += cur % 10
			cur /= 10
		}
		m[sum]++
	}
	max := 0
	count := 1
	for _, c := range m {
		if c == max {
			count++
		} else if c > max {
			max = c
			count = 1
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
