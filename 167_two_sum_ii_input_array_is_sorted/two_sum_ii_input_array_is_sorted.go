package two_sum_ii_input_array_is_sorted

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

// level: 1
// time: O(n*log(n)) 4ms 94.87%
// space: O(1) 3M

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
			continue
		}
		if numbers[i]+numbers[j] < target {
			i++
			continue
		}
		if numbers[i]+numbers[j] == target {
			break
		}
	}
	return []int{i + 1, j + 1}
}

// leetcode submit region end(Prohibit modification and deletion)
