package average_salary_excluding_the_minimum_and_maximum_salary

import "math"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M 7.69%

// leetcode submit region begin(Prohibit modification and deletion)
func average(salary []int) float64 {
	var max, min int = 0, 10e6
	sum := 0
	for _, s := range salary {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
		sum += s
	}
	sum -= max
	sum -= min

	f := float64(sum) / float64(len(salary)-2)
	return math.Trunc(f*1e5) / 1e5
}

// leetcode submit region end(Prohibit modification and deletion)
