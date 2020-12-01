package two_sum

// https://leetcode.com/problems/two-sum

// level: 1
// time: O(n) 4ms 95.58%
// space: O(n) 3.2M 79.77%

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)

	for i, num := range nums {
		if sumMap[num] > 0 {
			return []int{sumMap[num] - 1, i}
		} else {
			sumMap[target-num] = i + 1
		}
	}

	return []int{}
}

// leetcode submit region end(Prohibit modification and deletion)
