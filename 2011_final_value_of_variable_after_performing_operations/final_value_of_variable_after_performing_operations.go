package final_value_of_variable_after_performing_operations

// https://leetcode.com/problems/final-value-of-variable-after-performing-operations

// level: 1
// time: O(n) 4ms 100%
// space: O(1) 3.6M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, str := range operations {
		if str == "++X" || str == "X++" {
			x++
		} else {
			x--
		}
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
