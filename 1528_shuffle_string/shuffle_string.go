package shuffle_string

// https://leetcode.com/problems/shuffle-string

// level: 1
// time: O(n) 4ms 87.50%
// space: O(n) 3.4M 51.95%

//leetcode submit region begin(Prohibit modification and deletion)
func restoreString(s string, indices []int) string {
	chars := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		chars[indices[i]] = s[i]
	}
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)
