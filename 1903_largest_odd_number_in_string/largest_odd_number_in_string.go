package largest_odd_number_in_string

// https://leetcode.com/problems/largest-odd-number-in-string

// level: 1
// time: O(n) 4ms 97.06%
// space: O(1) 6.6M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func largestOddNumber(num string) string {
	n := len(num)
	i := n - 1
	for ; i >= 0; i-- {
		char := num[i] - '0'
		if char%2 != 0 {
			break
		}
	}
	return num[:i+1]
}

//leetcode submit region end(Prohibit modification and deletion)
